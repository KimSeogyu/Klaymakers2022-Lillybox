package http

import (
	"lillybox-backend/internal/database"
	"log"
)

// FatalWithError ...
func FatalWithError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// ParseLillyVidMany ...
func ParseLillyVidMany(vid []*database.Vids) []LillyVideo {
	var result []LillyVideo
	for _, v := range vid {
		result = append(result, ParseLillyVideo(v))
	}
	return result
}

// ParseLillyComment ...
func ParseLillyComment(comment database.Comments) LillyComment {
	return LillyComment{
		ID:          comment.ID,
		Nickname:    comment.Nickname,
		Description: comment.Description,
		CretaedAt:   comment.CreatedAt,
		UpdatedAt:   comment.UpdatedAt,
	}
}

// ParseLillyCommentMany ...
func ParseLillyCommentMany(comments []database.Comments) []LillyComment {
	var result []LillyComment
	for _, v := range comments {
		result = append(result, ParseLillyComment(v))
	}
	return result
}

// ParseLillyVideo ...
func ParseLillyVideo(vid *database.Vids) LillyVideo {
	return LillyVideo{
		ID:           vid.ID,
		CID:          vid.ContentID,
		Nickname:     vid.Nickname,
		Name:         vid.VideoName,
		Description:  vid.Description,
		Type:         "videos",
		Categories:   ParseCategories(vid.Category),
		CreatedAt:    vid.CreatedAt,
		VideoURI:     vid.VideoURI,
		ThumbnailURI: vid.ThumbnailURI,
		Views:        int(vid.VidViews.Count),
	}
}

// ParseCategories ...
func ParseCategories(categories []*database.Category) []string {
	result := []string{}
	for _, v := range categories {
		result = append(result, v.Name)
	}
	return result
}

// KeyCheck ...
func KeyCheck(data any, key string) bool {
	if _, ok := data.(map[string]string)[key]; ok {
		return true
	}
	return false
}
