package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
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

type AlbumsG struct {
	AlbumType string `json:"album_type"`
	Artists   []struct {
		ExternalUrls struct {
			Spotify string `json:"spotify"`
		} `json:"external_urls"`
		Href string `json:"href"`
		ID   string `json:"id"`
		Name string `json:"name"`
		Type string `json:"type"`
		URI  string `json:"uri"`
	} `json:"artists"`
	AvailableMarkets []string `json:"available_markets"`
	Copyrights       []struct {
		Text string `json:"text"`
		Type string `json:"type"`
	} `json:"copyrights"`
	ExternalIds struct {
		Upc string `json:"upc"`
	} `json:"external_ids"`
	ExternalUrls struct {
		Spotify string `json:"spotify"`
	} `json:"external_urls"`
	Genres []interface{} `json:"genres"`
	Href   string        `json:"href"`
	ID     string        `json:"id"`
	Images []struct {
		Height int    `json:"height"`
		URL    string `json:"url"`
		Width  int    `json:"width"`
	} `json:"images"`
	Label                string `json:"label"`
	Name                 string `json:"name"`
	Popularity           int    `json:"popularity"`
	ReleaseDate          string `json:"release_date"`
	ReleaseDatePrecision string `json:"release_date_precision"`
	TotalTracks          int    `json:"total_tracks"`
	Tracks               struct {
		Href  string `json:"href"`
		Items []struct {
			Artists []struct {
				ExternalUrls struct {
					Spotify string `json:"spotify"`
				} `json:"external_urls"`
				Href string `json:"href"`
				ID   string `json:"id"`
				Name string `json:"name"`
				Type string `json:"type"`
				URI  string `json:"uri"`
			} `json:"artists"`
			AvailableMarkets []string `json:"available_markets"`
			DiscNumber       int      `json:"disc_number"`
			DurationMs       int      `json:"duration_ms"`
			Explicit         bool     `json:"explicit"`
			ExternalUrls     struct {
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
		Limit    int         `json:"limit"`
		Next     interface{} `json:"next"`
		Offset   int         `json:"offset"`
		Previous interface{} `json:"previous"`
		Total    int         `json:"total"`
	} `json:"tracks"`
	Type string `json:"type"`
	URI  string `json:"uri"`
}

type Playlist struct {
	Collaborative bool   `json:"collaborative"`
	Description   string `json:"description"`
	ExternalUrls  struct {
		Spotify string `json:"spotify"`
	} `json:"external_urls"`
	Followers struct {
		Href  interface{} `json:"href"`
		Total int         `json:"total"`
	} `json:"followers"`
	Href   string `json:"href"`
	ID     string `json:"id"`
	Images []struct {
		Height interface{} `json:"height"`
		URL    string      `json:"url"`
		Width  interface{} `json:"width"`
	} `json:"images"`
	Name  string `json:"name"`
	Owner struct {
		DisplayName  string `json:"display_name"`
		ExternalUrls struct {
			Spotify string `json:"spotify"`
		} `json:"external_urls"`
		Href string `json:"href"`
		ID   string `json:"id"`
		Type string `json:"type"`
		URI  string `json:"uri"`
	} `json:"owner"`
	PrimaryColor string `json:"primary_color"`
	Public       bool   `json:"public"`
	SnapshotID   string `json:"snapshot_id"`
	Tracks       struct {
		Href  string `json:"href"`
		Items []struct {
			AddedAt time.Time `json:"added_at"`
			AddedBy struct {
				ExternalUrls struct {
					Spotify string `json:"spotify"`
				} `json:"external_urls"`
				Href string `json:"href"`
				ID   string `json:"id"`
				Type string `json:"type"`
				URI  string `json:"uri"`
			} `json:"added_by"`
			IsLocal      bool        `json:"is_local"`
			PrimaryColor interface{} `json:"primary_color"`
			Track        struct {
				Album struct {
					AlbumType string `json:"album_type"`
					Artists   []struct {
						ExternalUrls struct {
							Spotify string `json:"spotify"`
						} `json:"external_urls"`
						Href string `json:"href"`
						ID   string `json:"id"`
						Name string `json:"name"`
						Type string `json:"type"`
						URI  string `json:"uri"`
					} `json:"artists"`
					AvailableMarkets []string `json:"available_markets"`
					ExternalUrls     struct {
						Spotify string `json:"spotify"`
					} `json:"external_urls"`
					Href   string `json:"href"`
					ID     string `json:"id"`
					Images []struct {
						Height int    `json:"height"`
						URL    string `json:"url"`
						Width  int    `json:"width"`
					} `json:"images"`
					Name                 string `json:"name"`
					ReleaseDate          string `json:"release_date"`
					ReleaseDatePrecision string `json:"release_date_precision"`
					TotalTracks          int    `json:"total_tracks"`
					Type                 string `json:"type"`
					URI                  string `json:"uri"`
				} `json:"album"`
				Artists []struct {
					ExternalUrls struct {
						Spotify string `json:"spotify"`
					} `json:"external_urls"`
					Href string `json:"href"`
					ID   string `json:"id"`
					Name string `json:"name"`
					Type string `json:"type"`
					URI  string `json:"uri"`
				} `json:"artists"`
				AvailableMarkets []string `json:"available_markets"`
				DiscNumber       int      `json:"disc_number"`
				DurationMs       int      `json:"duration_ms"`
				Episode          bool     `json:"episode"`
				Explicit         bool     `json:"explicit"`
				ExternalIds      struct {
					Isrc string `json:"isrc"`
				} `json:"external_ids"`
				ExternalUrls struct {
					Spotify string `json:"spotify"`
				} `json:"external_urls"`
				Href        string `json:"href"`
				ID          string `json:"id"`
				IsLocal     bool   `json:"is_local"`
				Name        string `json:"name"`
				Popularity  int    `json:"popularity"`
				PreviewURL  string `json:"preview_url"`
				Track       bool   `json:"track"`
				TrackNumber int    `json:"track_number"`
				Type        string `json:"type"`
				URI         string `json:"uri"`
			} `json:"track"`
			VideoThumbnail struct {
				URL interface{} `json:"url"`
			} `json:"video_thumbnail"`
		} `json:"items"`
		Limit    int         `json:"limit"`
		Next     interface{} `json:"next"`
		Offset   int         `json:"offset"`
		Previous interface{} `json:"previous"`
		Total    int         `json:"total"`
	} `json:"tracks"`
	Type string `json:"type"`
	URI  string `json:"uri"`
}

type AlbumR struct {
	Albums struct {
		Href  string `json:"href"`
		Items []struct {
			AlbumType string `json:"album_type"`
			Artists   []struct {
				ExternalUrls struct {
					Spotify string `json:"spotify"`
				} `json:"external_urls"`
				Href string `json:"href"`
				ID   string `json:"id"`
				Name string `json:"name"`
				Type string `json:"type"`
				URI  string `json:"uri"`
			} `json:"artists"`
			AvailableMarkets []string `json:"available_markets"`
			ExternalUrls     struct {
				Spotify string `json:"spotify"`
			} `json:"external_urls"`
			Href   string `json:"href"`
			ID     string `json:"id"`
			Images []struct {
				Height int    `json:"height"`
				URL    string `json:"url"`
				Width  int    `json:"width"`
			} `json:"images"`
			Name                 string `json:"name"`
			ReleaseDate          string `json:"release_date"`
			ReleaseDatePrecision string `json:"release_date_precision"`
			TotalTracks          int    `json:"total_tracks"`
			Type                 string `json:"type"`
			URI                  string `json:"uri"`
		} `json:"items"`
		Limit    int         `json:"limit"`
		Next     string      `json:"next"`
		Offset   int         `json:"offset"`
		Previous interface{} `json:"previous"`
		Total    int         `json:"total"`
	} `json:"albums"`
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
			External_urls struct {
				Spotify string `json:"spotify"`
			} `json:"external_urls"`

			Followers struct {
				Total int `json:"total"`
			} `json:"followers"`
			Genres []string `json:"genres"`
			Id     string   `json:"id"`
			Image  []struct {
				URL string `json:"url"`
			} `json:"images"`
			Href string `json:"href"`
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

type Favorites struct {
	IDs []string `json:"ids"`
}

type Albums struct {
	AlbumType string `json:"album_type"`
	Artists   []struct {
		ExternalUrls struct {
			Spotify string `json:"spotify"`
		} `json:"external_urls"`
		Href string `json:"href"`
		ID   string `json:"id"`
		Name string `json:"name"`
		Type string `json:"type"`
		URI  string `json:"uri"`
	} `json:"artists"`
	AvailableMarkets []string `json:"available_markets"`
	Copyrights       []struct {
		Text string `json:"text"`
		Type string `json:"type"`
	} `json:"copyrights"`
	ExternalIds struct {
		Upc string `json:"upc"`
	} `json:"external_ids"`
	ExternalUrls struct {
		Spotify string `json:"spotify"`
	} `json:"external_urls"`
	Genres []interface{} `json:"genres"`
	Href   string        `json:"href"`
	ID     string        `json:"id"`
	Images []struct {
		Height int    `json:"height"`
		URL    string `json:"url"`
		Width  int    `json:"width"`
	} `json:"images"`
	Label                string `json:"label"`
	Name                 string `json:"name"`
	Popularity           int    `json:"popularity"`
	ReleaseDate          string `json:"release_date"`
	ReleaseDatePrecision string `json:"release_date_precision"`
	TotalTracks          int    `json:"total_tracks"`
	Tracks               struct {
		Href  string `json:"href"`
		Items []struct {
			Artists []struct {
				ExternalUrls struct {
					Spotify string `json:"spotify"`
				} `json:"external_urls"`
				Href string `json:"href"`
				ID   string `json:"id"`
				Name string `json:"name"`
				Type string `json:"type"`
				URI  string `json:"uri"`
			} `json:"artists"`
			AvailableMarkets []string `json:"available_markets"`
			DiscNumber       int      `json:"disc_number"`
			DurationMs       int      `json:"duration_ms"`
			Explicit         bool     `json:"explicit"`
			ExternalUrls     struct {
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
		Limit    int         `json:"limit"`
		Next     interface{} `json:"next"`
		Offset   int         `json:"offset"`
		Previous interface{} `json:"previous"`
		Total    int         `json:"total"`
	} `json:"tracks"`
	Type string `json:"type"`
	URI  string `json:"uri"`
}

type PopularArtists struct {
	Artists []Artist `json:"artists"`
}

type Artist struct {
	Followers struct {
		Total int `json:"total"`
	} `json:"followers"`
	Images []struct {
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

func getbyArtists(id_track string, token string) (TrackInfo, error) {
	url := fmt.Sprintf("https://api.spotify.com/v1/albums/%v/tracks", id_track)
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
	fmt.Println(resultList)
	return resultList, nil
}

func getArtists(id_artist string, token string) (Artist, error) {
	url := fmt.Sprintf("https://api.spotify.com/v1/search?q=genre:pop&type=artist")
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

func getPopPlaylists(id_artist string, token string) (Playlist, error) {
	url := fmt.Sprintf("https://api.spotify.com/v1/playlists/%v", id_artist)
	res, resErr := requestSpotifyAPI(url, token)
	if resErr != nil {
		return Playlist{}, fmt.Errorf("failed to fetch playlist: %v", resErr) // Remarquez Playlist{} au lieu de nil
	}
	defer res.Body.Close()

	var response Playlist

	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return Playlist{}, fmt.Errorf("failed to decode playlist response: %v", err) // Remarquez Playlist{} au lieu de nil
	}
	return response, nil
}

func getRapPlaylists(id_artist string, token string) (Playlist, error) {
	url := fmt.Sprintf("https://api.spotify.com/v1/playlists/%v", id_artist)
	res, resErr := requestSpotifyAPI(url, token)
	if resErr != nil {
		return Playlist{}, fmt.Errorf("failed to fetch playlist: %v", resErr) // Remarquez Playlist{} au lieu de nil
	}
	defer res.Body.Close()

	var response Playlist

	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return Playlist{}, fmt.Errorf("failed to decode playlist response: %v", err) // Remarquez Playlist{} au lieu de nil
	}
	return response, nil
}

func getPhonkPlaylists(id_artist string, token string) (Playlist, error) {
	url := fmt.Sprintf("https://api.spotify.com/v1/playlists/%v", id_artist)
	res, resErr := requestSpotifyAPI(url, token)
	if resErr != nil {
		return Playlist{}, fmt.Errorf("failed to fetch playlist: %v", resErr) // Remarquez Playlist{} au lieu de nil
	}
	defer res.Body.Close()

	var response Playlist

	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return Playlist{}, fmt.Errorf("failed to decode playlist response: %v", err) // Remarquez Playlist{} au lieu de nil
	}
	return response, nil
}

func getAlbums(id_jul string, token string) (Album, error) {
	url := fmt.Sprintf("https://api.spotify.com/v1/artists/%v/albums", id_jul)
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

func getPopularAlbums(id_artist string, token string) (Albums, error) {
	url := fmt.Sprintf("https://api.spotify.com/v1/artists/%v/albums", id_artist)
	res, resErr := requestSpotifyAPI(url, token)
	if resErr != nil {
		return Albums{}, fmt.Errorf(resErr.Error())

	}
	defer res.Body.Close()
	var resultList Albums
	errDecode := json.NewDecoder(res.Body).Decode(&resultList)
	if errDecode != nil {
		return Albums{}, fmt.Errorf(errDecode.Error())
	}
	return resultList, nil

}

func getAlbumsG(id string, token string) (AlbumsG, error) {
	url := fmt.Sprintf("https://api.spotify.com/v1/albums/%v", id)
	res, resErr := requestSpotifyAPI(url, token)
	if resErr != nil {
		return AlbumsG{}, fmt.Errorf(resErr.Error())

	}
	defer res.Body.Close()
	var resultList AlbumsG
	errDecode := json.NewDecoder(res.Body).Decode(&resultList)
	if errDecode != nil {
		return AlbumsG{}, fmt.Errorf(errDecode.Error())
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

func AddToFavorites(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id_artist := r.FormValue("id_artist")
	fmt.Println(id_artist)

	var favs Favorites
	file, err := ioutil.ReadFile("favorites.json")
	if err != nil {
		// Gérer l'erreur ou créer le fichier s'il n'existe pas
	}
	json.Unmarshal(file, &favs)

	favs.IDs = append(favs.IDs, id_artist)

	// Écrire le fichier JSON mis à jour
	updatedFavs, err := json.Marshal(favs)
	if err != nil {
		// Gérer l'erreur
	}

	if err != nil && !os.IsNotExist(err) {
		http.Error(w, "Erreur lors de la lecture des favoris", http.StatusInternalServerError)
		return
	} else if os.IsNotExist(err) {
		favs = Favorites{IDs: []string{}}
	}

	if err := ioutil.WriteFile("favorites.json", updatedFavs, 0644); err != nil {
		http.Error(w, "Erreur lors de la sauvegarde des favoris", http.StatusInternalServerError)
		return
	}

	// Redirection vers la page des favoris
	http.Redirect(w, r, "/favorites", http.StatusSeeOther)
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

func getReleased(token string) (AlbumR, error) {
	url := fmt.Sprintf("https://api.spotify.com/v1/browse/new-releases")
	res, resErr := requestSpotifyAPI(url, token)
	if resErr != nil {
		println("error")
		return AlbumR{}, fmt.Errorf(resErr.Error())

	}
	defer res.Body.Close()
	var resultList AlbumR
	println(json.NewDecoder(res.Body))
	errDecode := json.NewDecoder(res.Body).Decode(&resultList)
	if errDecode != nil {
		fmt.Println("Decode error")
		return AlbumR{}, fmt.Errorf(errDecode.Error())
	}
	return resultList, nil
}

func loadAlbumData(filename string) (*Album, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var data Album
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func showAlbumsHandler(w http.ResponseWriter, r *http.Request) {

	albumData, err := loadAlbumData("favorites.json")
	if err != nil {
		http.Error(w, "Failed to load album data", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("/templates/favorites.html")
	if err != nil {
		http.Error(w, "Failed to load template", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, albumData)
	if err != nil {
		http.Error(w, "Failed to execute template", http.StatusInternalServerError)
	}
}

func main() {
	var err error
	accessToken, err = tokenaccess()
	if err != nil {
		log.Fatalf("Failed to retrieve access token: %v", err)
	}
	data, errData := getAlbumsG("4VL5XwfATZuAVTW471Wpro", accessToken)
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
		artist, _ := getAlbums("3IW7ScrzXmPvZhB27hmfgy", accessToken)
		if err := tmpl.ExecuteTemplate(w, "home", artist); err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			log.Printf("Error executing template: %v", err)
			return
		}
	})

	// Définition du gestionnaire pour la racine du serveur
	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		query := r.FormValue("search")
		fmt.Println(query)
		artist, _ := searchArtists(query, accessToken)
		fmt.Println(artist)
		if err := tmpl.ExecuteTemplate(w, "search", artist); err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			log.Printf("Error executing template: %v", err)
			return
		}
	})

	/////playlist fr
	/////Page playlist
	http.HandleFunc("/rapfr", func(w http.ResponseWriter, r *http.Request) {
		id_artist := r.FormValue("id")
		fmt.Println(id_artist)

		artist, err := getRapPlaylists("37i9dQZF1DX1X23oiQRTB5", accessToken)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			log.Printf("Error getting artist data: %v", err)
			return
		}
		fmt.Println(artist) // Vérifiez les données récupérées de l'API Spotify

		if err := tmpl.ExecuteTemplate(w, "rapfr", artist); err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			log.Printf("Error executing template: %v", err)
			return
		}
	})

	///playlist phonk
	http.HandleFunc("/phonk", func(w http.ResponseWriter, r *http.Request) {
		id_artist := r.FormValue("id")
		fmt.Println(id_artist)

		artist, err := getPhonkPlaylists("37i9dQZF1DWWY64wDtewQt", accessToken)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			log.Printf("Error getting artist data: %v", err)
			return
		}
		fmt.Println(artist) // Vérifiez les données récupérées de l'API Spotify

		if err := tmpl.ExecuteTemplate(w, "phonk", artist); err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			log.Printf("Error executing template: %v", err)
			return
		}
	})
	///route favoris
	http.HandleFunc("/add-favorite", AddToFavorites)

	/////Page playlist pop
	http.HandleFunc("/pop", func(w http.ResponseWriter, r *http.Request) {
		id_artist := r.FormValue("id")
		fmt.Println(id_artist)

		artist, err := getPopPlaylists("0chdKQr18NN9WRI355V8BN", accessToken)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			log.Printf("Error getting artist data: %v", err)
			return
		}
		fmt.Println(artist) // Vérifiez les données récupérées de l'API Spotify

		if err := tmpl.ExecuteTemplate(w, "pop", artist); err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			log.Printf("Error executing template: %v", err)
			return
		}
	})

	http.HandleFunc("/favorites", func(w http.ResponseWriter, r *http.Request) {
		println("Favorites")
		var favs Favorites
		file, err := ioutil.ReadFile("favorites.json")
		if err != nil {
			// Gérer l'erreur ou créer le fichier s'il n'existe pas
			favs = Favorites{IDs: []string{}}
		}
		json.Unmarshal(file, &favs)

		var artists []AlbumsG
		for _, id := range favs.IDs {
			println(id)
			artist, err := getAlbumsG(id, accessToken)
			if err != nil {
				// Gérer l'erreur
				println("aaaaa")
			}
			artists = append(artists, artist)
		}
		fmt.Println(artists)
		tmpl.ExecuteTemplate(w, "favorites", artists)
	})

	http.HandleFunc("/release", func(w http.ResponseWriter, r *http.Request) {
		println("Release")
		artist, _ := getReleased(accessToken)
		fmt.Println(artist)
		if err := tmpl.ExecuteTemplate(w, "release", artist); err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			log.Printf("Error executing template: %v", err)
			return
		}
	})

	// Définition des gestionnaires de fichiers statiques
	css := http.FileServer(http.Dir("./styles"))
	http.Handle("/static/", http.StripPrefix("/static/", css))

	// Démarrage du serveur HTTP
	fmt.Println("Serveur lancé sur http://localhost:8089")
	log.Fatal(http.ListenAndServe(":8089", nil))
}
