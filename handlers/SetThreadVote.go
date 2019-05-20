package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"technopark_db/models"
)

func (h *Handler) SetThreadVote(c *gin.Context) {
	var vote models.Vote
	_ = json.NewDecoder(c.Request.Body).Decode(&vote)
	slug_or_id := c.Param("slug_or_id")
	id, err := strconv.Atoi(slug_or_id)
	if err != nil {
		id, err = h.Agregator.GetThreadVotesBySlug(slug_or_id)
		if err != nil {
			fmt.Println(err)
			c.JSON(404, err)
			return
		}
	}
	vote.Id = id
	oldVote, err := h.Agregator.GetVote(vote.Nickname, id)
	var thread models.Thread
	if err != nil {
		if err.Error() == "no rows in result set" {
			err = h.Agregator.InsertVote(vote)
			if err != nil {
				fmt.Println(err)
				c.JSON(404, err)
				return
			}
			thread, err = h.Agregator.UpdateThreadVote(vote.Voice, vote.Id)
			if err != nil {
				fmt.Println(err)
				c.JSON(404, err)
				return
			}
			c.JSON(200, thread)
			return
		}
	}
	thread, err = h.Agregator.GetThreadById(vote.Id)
	if oldVote.Voice == vote.Voice {
		thread, err = h.Agregator.GetThreadById(vote.Id)
		if err != nil {
			fmt.Println(err)
			c.JSON(404, err)
			return
		}
	} else {
		err = h.Agregator.UpdateVote(vote)
		if err != nil {
			fmt.Println(err)
			c.JSON(404, err)
			return
		}
		thread, err = h.Agregator.UpdateThreadVote(vote.Voice*2, vote.Id)
		if err != nil {
			fmt.Println(err)
			c.JSON(404, err)
			return
		}
	}
	c.JSON(200, thread)
}
