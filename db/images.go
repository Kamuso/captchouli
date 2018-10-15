package db

import (
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/bakape/captchouli/common"
)

type Image struct {
	Source common.DataSource
	MD5    [16]byte
	Tags   []string
}

// Return, if file is not already registered in the DB as valid thumbnail or in
// a blacklist
func IsInDatabase(md5 [16]byte) (is bool, err error) {
	dbMu.RLock()
	defer dbMu.RUnlock()

	err = sq.Select("1").
		From("images").
		Where("hash = ?", md5[:]).
		Scan(&is)
	if err == sql.ErrNoRows {
		err = nil
	}
	return
}

// Write image to database
func InsertImage(img Image) (err error) {
	dbMu.Lock()
	defer dbMu.Unlock()

	return InTransaction(func(tx *sql.Tx) (err error) {
		r, err := withTransaction(tx, insertImage(img.MD5)).Exec()
		if err != nil {
			return
		}
		id, err := r.LastInsertId()
		if err != nil {
			return
		}

		q, err := tx.Prepare(
			`insert into image_tags (image_id, tag, source)
			values(?, ?, ?)`)
		if err != nil {
			return
		}
		for _, t := range img.Tags {
			_, err = q.Exec(id, t, img.Source)
			if err != nil {
				return
			}
		}
		return
	})
}

func insertImage(hash [16]byte) squirrel.InsertBuilder {
	return sq.
		Insert("images").
		Columns("hash").
		Values(hash[:])
}

// Add image to blacklist so that it is not fetched again
func BlacklistImage(hash [16]byte) (err error) {
	dbMu.Lock()
	defer dbMu.Unlock()

	_, err = insertImage(hash).Exec()
	return
}

// Return count of images matching selectors
func ImageCount(tag string, src common.DataSource) (n int, err error) {
	dbMu.RLock()
	defer dbMu.RUnlock()

	err = sq.Select("count(*)").
		From("image_tags").
		Where(squirrel.Eq{
			"tag":    tag,
			"source": src,
		}).
		Scan(&n)
	return
}
