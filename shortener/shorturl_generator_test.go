package shortener

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const UserId = "e0dba740-fc4b-4977-872c-d360239e6b1a"

func TestShortLinkGenerator(t *testing.T) {
	initialURL_1 := "https://www.guru3d.com/news-story/spotted-ryzen-threadripper-pro-3995wx-processor-with-8-channel-ddr4,2.html"
	shortURL_1 := GenerateShortLink(initialURL_1, UserId)

	initialURL_2 := "https://www.eddywm.com/lets-build-a-url-shortener-in-go-with-redis-part-2-storage-layer/"
	shortURL_2 := GenerateShortLink(initialURL_2, UserId)

	initialURL_3 := "https://spectrum.ieee.org/automaton/robotics/home-robots/hello-robots-stretch-mobile-manipulator"
	shortURL_3 := GenerateShortLink(initialURL_3, UserId)

	assert.Equal(t, shortURL_1, "jTa4L57P")
	assert.Equal(t, shortURL_2, "d66yfx7N")
	assert.Equal(t, shortURL_3, "dhZTayYQ")
}
