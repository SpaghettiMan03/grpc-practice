package handler

import (
	"bytes"
	"grpc-practice/gen/api"
	"io"
	"net/http"
	"sync"

	"github.com/google/uuid"
)

type ImageUploadHandler struct {
	sync.Mutex
	file map[string][]byte
}

func (h *ImageUploadHandler) Upload(stream api.ImageUploadService_UploadServer) error {
	req, err := stream.Recv()
	if err != nil {
		return err
	}

	meta := req.GetFileMeta()
	fileName := meta.Filename

	u, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	uuid := u.String()

	buf := &bytes.Buffer{}

	for {
		recv, err := stream.Recv()

		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		chunk := recv.GetData()
		_, err = buf.Write(chunk)
		if err != nil {
			return err
		}
	}

	data := buf.Bytes()
	mimeType := http.DetectContentType(data)

	h.file[fileName] = data

	err = stream.SendAndClose(&api.ImageUploadResponse{
		Uuid:        uuid,
		Size:        int32(len(data)),
		Filename:    fileName,
		ContentType: mimeType,
	})

	return err
}

func NewImageUploadHandler() *ImageUploadHandler {
	return &ImageUploadHandler{
		file: make(map[string][]byte),
	}
}
