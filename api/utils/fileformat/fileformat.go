package fileformat

import (
	"path"
	"strings"

	"github.com/twinj/uuid"
)

/*
A user will need to update his profile(including adding an image)
when he does, we will need to make sure that we image added has a unique name.

In the utils directory(path: /forum-backend/api/utils),
create the fileformat directory.
 */

func UniqueFormat(fn string) string {
	//path.Ext() get the extension of the file
	fileName := strings.TrimSuffix(fn, path.Ext(fn))
	extension := path.Ext(fn)
	u := uuid.NewV4()
	newFileName := fileName + "-" + u.String() + extension

	return newFileName

}