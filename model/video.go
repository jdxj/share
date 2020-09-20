package model

import "fmt"

const (
	TNVideo = "videos"

	CountPage = 20
)

func GetVideos(page int) ([]*Video, error) {
	query := fmt.Sprintf(`select id,title,path,user_id
from %s
order by id limit ?,?`, TNVideo)
	rows, err := mysql.Query(query, page*CountPage, CountPage)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var vides []*Video
	for rows.Next() {
		v := &Video{}
		err := rows.Scan(&v.ID, &v.Title, &v.Path, &v.UserID)
		if err != nil {
			return nil, err
		}
		vides = append(vides, v)
	}
	return vides, nil
}

func GetVideoByID(id int) (*Video, error) {
	query := fmt.Sprintf(`select id,title,path,user_id
from %s where id=?`, TNVideo)
	row := mysql.QueryRow(query, id)
	v := &Video{}
	return v, row.Scan(&v.ID, &v.Title, &v.Path, &v.UserID)
}

type Video struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Path   string `json:"path"`
	UserID int    `json:"user_id"`
}

func (v *Video) Insert() error {
	query := fmt.Sprintf(`insert into %s (title,path,user_id)
values (?,?,?)`, TNVideo)
	_, err := mysql.Exec(query, v.Title, v.Path, v.UserID)
	return err
}
