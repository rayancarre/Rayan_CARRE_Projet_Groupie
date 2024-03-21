package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

var (
	clientID     = "b46d215af0cc40d28a63b64d810c2e2c"
	clientSecret = "113870a84d914a2393a8e31719eb241c"
	accessToken  string
)

type ExternalURLsSearch struct {
	Spotify string `json:"spotify"`
}

// FollowersSearch représente les abonnés.
type FollowersSearch struct {
	Href  interface{} `json:"href"`
	Total int         `json:"total"`
}

// ImageSearch représente une image.
type ImageSearch struct {
	URL    string `json:"url"`
	Height int    `json:"height"`
	Width  int    `json:"width"`
}

// ArtistSearch représente un artiste.
type ArtistSearch struct {
	ExternalURLs ExternalURLsSearch `json:"external_urls"`
	Followers    FollowersSearch    `json:"followers"`
	Genres       []string           `json:"genres"`
	Href         string             `json:"href"`
	ID           string             `json:"id"`
	Images       []ImageSearch      `json:"images"`
	Name         string             `json:"name"`
	Popularity   int                `json:"popularity"`
	Type         string             `json:"type"`
	URI          string             `json:"uri"`
}

// ArtistsResponseSearch représente la réponse de l'API Spotify pour les artistes.
type ArtistsResponseSearch struct {
	Artists struct {
		Href     string         `json:"href"`
		Limit    int            `json:"limit"`
		Next     string         `json:"next"`
		Offset   int            `json:"offset"`
		Previous string         `json:"previous"`
		Total    int            `json:"total"`
		Items    []ArtistSearch `json:"items"`
	} `json:"artists"`
}

type Image struct {
	URL string `json:"url"`
}

type TrackInfo struct {
	Items []struct {
		Artists []struct {
			ExternalUrls struct {
				Spotify string `json:"spotify"`
			} `json:"external_urls"`
			Href string `json:"href"`
			ID   string `json:"id"`
			Name string `json:"name"`
			URI  string `json:"uri"`
		} `json:"artists"`
		DiscNumber   int  `json:"disc_number"`
		DurationMs   int  `json:"duration_ms"`
		Explicit     bool `json:"explicit"`
		ExternalUrls struct {
			Spotify string `json:"spotify"`
		} `json:"external_urls"`
		Href        string `json:"href"`
		ID          string `json:"id"`
		IsLocal     bool   `json:"is_local"`
		Name        string `json:"name"`
		PreviewURL  string `json:"preview_url"`
		TrackNumber int    `json:"track_number"`
		Type        string `json:"type"`
		URI         string `json:"uri"`
	} `json:"items"`
}

type SearchResults struct {
	Albums []Album `json:"items"`
}

func requestSpotifyAPI(url, token string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+token)
	client := &http.Client{}
	return client.Do(req)
}

type SearchArtists struct {
	Artists struct {
		Items []struct {
			Followers struct {
				Total int `json:"total"`
			} `json:"followers"`
			Genres []string `json:"genres"`
			Id     string   `json:"id"`
			Image  []struct {
				URL string `json:"url"`
			} `json:"images"`
			Name string `json:"name"`
		} `json:"items"`
	} `json:"artists"`
}

type Album struct {
	Items []struct {
		Artist []struct {
			Name          string `json:"name"`
			Id            string `json:"id"`
			External_urls struct {
				Spotify string `json:" spotify "`
			} `json:"external_urls"`
		} `json:"artists"`
		External_urls struct {
			Spotify string `json:" spotify "`
		} `json:"external_urls"`
		Id     string `json:"id"`
		Images []struct {
			Url string `json:"url"`
		} `json:"images"`
		Name         string `json:"name"`
		Release_date string `json:"release_date"`
		Total_tracks int    `json:"total_tracks"`
	} `json:"items"`
}

type Artist struct {
	Followers struct {
		Total int `json:"total"`
	} `json:"followers"`
	Image []struct {
		URL string `json:"url"`
	} `json:"images"`
	Name       string `json:"name"`
	Popularity int    `json:"popularity"`
}

func searchArtists(q string, token string) (SearchArtists, error) {
	url := fmt.Sprintf("https://api.spotify.com/v1/search?q=%v&type=artist&market=FR", q)
	res, resErr := requestSpotifyAPI(url, token)
	if resErr != nil {
		return SearchArtists{}, fmt.Errorf(resErr.Error())

	}
	defer res.Body.Close()
	var resultList SearchArtists
	errDecode := json.NewDecoder(res.Body).Decode(&resultList)
	if errDecode != nil {
		return SearchArtists{}, fmt.Errorf(errDecode.Error())
	}
	return resultList, nil
}

func getArtists(id_artist string, token string) (Artist, error) {
	url := fmt.Sprintf("https://api.spotify.com/v1/artists/%v", id_artist)
	res, resErr := requestSpotifyAPI(url, token)
	if resErr != nil {
		return Artist{}, fmt.Errorf(resErr.Error())
	}

	defer res.Body.Close()
	var resultList Artist
	errDecode := json.NewDecoder(res.Body).Decode(&resultList)
	if errDecode != nil {
		return Artist{}, fmt.Errorf(errDecode.Error())
	}
	return resultList, nil
}

func getAlbums(id_artist string, token string) (Album, error) {
	url := fmt.Sprintf("https://api.spotify.com/v1/artists/%v/albums", id_artist)
	res, resErr := requestSpotifyAPI(url, token)
	if resErr != nil {
		return Album{}, fmt.Errorf(resErr.Error())

	}
	defer res.Body.Close()
	var resultList Album
	errDecode := json.NewDecoder(res.Body).Decode(&resultList)
	if errDecode != nil {
		return Album{}, fmt.Errorf(errDecode.Error())
	}
	return resultList, nil

}

func getMusics(id_music string, token string) (TrackInfo, error) {
	url := fmt.Sprintf("https://api.spotify.com/v1/albums/%v/tracks", id_music)
	res, resErr := requestSpotifyAPI(url, token)
	if resErr != nil {
		return TrackInfo{}, fmt.Errorf(resErr.Error())

	}
	defer res.Body.Close()
	var resultList TrackInfo
	errDecode := json.NewDecoder(res.Body).Decode(&resultList)
	if errDecode != nil {
		return TrackInfo{}, fmt.Errorf(errDecode.Error())
	}
	return resultList, nil

}

func getRecentAlbums(token string) ([]Album, error) {
	id_artist := "3IW7ScrzXmPvZhB27hmfgy"
	url := "https://api.spotify.com/v1/artists/" + id_artist + "/albums?include_groups=single&limit=50"
	resp, err := requestSpotifyAPI(url, token)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	fmt.Println(resp.StatusCode)
	var searchResults SearchResults
	if err := json.NewDecoder(resp.Body).Decode(&searchResults); err != nil {
		return nil, err
	}

	return searchResults.Albums, nil
}

func tokenaccess() (string, error) {
	clientCreds := fmt.Sprintf("%s:%s", clientID, clientSecret)
	clientCredsB64 := base64.StdEncoding.EncodeToString([]byte(clientCreds))

	tokenURL := "https://accounts.spotify.com/api/token"
	tokenData := strings.NewReader("grant_type=client_credentials")

	req, err := http.NewRequest("POST", tokenURL, tokenData)
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Basic "+clientCredsB64)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("spotify API returned non-OK status: %d", resp.StatusCode)
	}

	var tokenResponse map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&tokenResponse); err != nil {
		return "", err
	}

	token, ok := tokenResponse["access_token"].(string)
	if !ok {
		return "", fmt.Errorf("invalid token format")
	}
	fmt.Println(token)
	return token, nil
}

func main() {
	var err error
	accessToken, err = tokenaccess()
	if err != nil {
		log.Fatalf("Failed to retrieve access token: %v", err)
	}
	data, errData := getMusics("49gSslDoncGfaxtZfsHyTA", accessToken)
	if errData != nil {
		fmt.Println(errData)
	}
	fmt.Println(data)
	tmpl, err := template.ParseGlob("./templates/*.html")
	if err != nil {
		log.Printf("Error parsing template: %v", err)
		return
	}

	// Définition du gestionnaire pour la racine du serveur
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		albums, _ := getRecentAlbums(accessToken)
		if err := tmpl.ExecuteTemplate(w, "home", albums); err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			log.Printf("Error executing template: %v", err)
			return
		}
	})

	// Définition des gestionnaires de fichiers statiques
	css := http.FileServer(http.Dir("./styles"))
	http.Handle("/static/", http.StripPrefix("/static/", css))

	// Démarrage du serveur HTTP
	fmt.Println("Serveur lancé sur http://localhost:8086")
	log.Fatal(http.ListenAndServe(":8086", nil))
}
