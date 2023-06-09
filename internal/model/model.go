package model

type ErrorResponse struct {
	ErrorMessage string
}

type User struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"Password" validate:"required"`
}

type CreateNewSessionRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"Password" validate:"required"`
}

type CreateNewSessionResponse struct {
	SerialId string `json:"sid"`
}

type Note struct {
	Id        uint32 `json:"id"`
	NoteEntry string `json:"note"`
}

type GetNotesRequest struct {
	SerialId string `json:"sid" validate:"required"`
}

type GetNotesResponse struct {
	Notes []Note `json:"notes"`
}

type CreateNoteRequest struct {
	SerialId string `json:"sid" validate:"required"`
	Note     string `json:"note" validate:"required"`
}

type CreateNoteResponse struct {
	Id uint32 `json:"id"`
}

type DeleteNoteRequest struct {
	SerialId string `json:"sid" validate:"required"`
	Id       uint32 `json:"id" validate:"required"`
}
