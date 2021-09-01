package mysql

import (
	"context"
	"database/sql"
	"time"

	"github.com/DTGlov/go-books/pkg/models"
)

type BookModel struct {
	DB *sql.DB
}

func (m *BookModel) Get(id int) (*models.Books,error){

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `SELECT id,title,description,author_id,publisher_id,created_at,updated_at FROM books WHERE id = ?`

	row := m.DB.QueryRowContext(ctx,stmt,id)

	 b := &models.Books{}
	err := row.Scan(
		&b.ID,
		&b.Title,
		&b.Description,
		&b.AuthorID,
		&b.PublisherID,
		&b.CreatedAt,
		&b.UpdatedAt,
	)
		if err != nil {
		return nil, err
	}
	

	//get authors 
	stmt = `select b.id,a.name,a.nationality,a.created_at,a.updated_at from books as b  join authors as a on b.author_id = a.id where b.id = ?`


	rows := m.DB.QueryRowContext(ctx,stmt,id)

	a:= models.Authors{}
	err= rows.Scan(
		&a.ID,
		&a.Name,
		&a.Nationality,
		&a.CreatedAt,
		&a.UpdatedAt,

	)
	if err !=nil{
		return nil,err
	}
	b.Author = a

	//get publishers
	stmt = `select b.id,p.name,p.publisher_id,p.created_at,p.updated_at from books as b left join publishers as p on b.publisher_id = p.publisher_id  where b.id = ?`


	rowss:= m.DB.QueryRowContext(ctx,stmt,id)

	p:= models.Publishers{}
	err = rowss.Scan(
		&p.ID,
		&p.Name,
		&b.PublisherID,
		&p.CreatedAt,
		&p.UpdatedAt,
	)
		if err !=nil{
		return nil,err
	}
	b.Publisher  = p
	

	return b,nil 
}

func (m *BookModel) GetAll()([]*models.Books,error){
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `SELECT id,title,description,author_id,publisher_id,created_at,updated_at FROM books`

	rows,err := m.DB.QueryContext(ctx,stmt)
	if err != nil {
		return  nil,err
	}
	defer rows.Close()

	books := []*models.Books{}
	
	for rows.Next(){
		b := &models.Books{}
		err = rows.Scan(
		&b.ID,
		&b.Title,
		&b.Description,
		&b.AuthorID,
		&b.PublisherID,
		&b.CreatedAt,
		&b.UpdatedAt,
		)
		if err !=nil{
			return nil,err
		}
		//Append it to the slice of books


		//get authors
		authorQuery := `select b.id,a.name,a.nationality,a.created_at,a.updated_at from books as b  join authors as a on b.author_id = a.id where b.id =?`

		authorRows,err :=m.DB.QueryContext(ctx,authorQuery,b.ID)
		if err != nil {
		return  nil,err
	}

	
		var authors = &models.Authors{}
		for authorRows.Next(){
			a:= &models.Authors{}
			err := authorRows.Scan(
			&a.ID,
			&a.Name,
			&a.Nationality,
			&a.CreatedAt,
			&a.UpdatedAt,
			)
			if err !=nil{
				return nil,err
			}
			authors = a
		}
		b.Author = *authors
		authorRows.Close()

		// get publishers if there are any
			pubQuery :=  `select b.id,p.name,p.publisher_id,p.created_at,p.updated_at from books as b left join publishers as p on b.publisher_id = p.publisher_id  where b.id = ?`
			pubRows,err := m.DB.QueryContext(ctx,pubQuery,b.ID)
			if err != nil {
		return  nil,err
			}

			var publishers = &models.Publishers{}
			for pubRows.Next(){
				p := &models.Publishers{}
				err :=pubRows.Scan(
				&p.ID,
				&p.Name,
				&b.PublisherID,
				&p.CreatedAt,
				&p.UpdatedAt,
				)
				if err !=nil{
				return nil,err
			}
			publishers =  p
			}
			b.Publisher = *publishers
			pubRows.Close()



		books = append(books, b)

	}
	  // When the rows.Next() loop has finished we call rows.Err() to retrieve any
    // error that was encountered during the iteration. It's important to
    // call this - don't assume that a successful iteration was completed
	if err = rows.Err();err !=nil{
		return nil,err
	}
	//if everything is ok then return the books slice
	return books,nil
}

func (m *BookModel) GetGenres()([]*models.Genre,error){
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `select id,genre_name from genre `

	rows,err := m.DB.QueryContext(ctx,stmt)
	if err !=nil {
		return  nil,err
	}
	defer rows.Close()

	genre := []*models.Genre{}

	for rows.Next(){
		g := &models.Genre{}
		err = rows.Scan(
			&g.ID,
			&g.GenreName,
		)
		if err !=nil{
			return nil,err
		}
		genre = append(genre, g)
	}
		if err = rows.Err();err !=nil{
		return nil,err
	}
	return genre,nil

}