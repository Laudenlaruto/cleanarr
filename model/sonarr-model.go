package model

import "time"

type SonarStatusResponse struct {
	Version           string    `json:"version"`
	Buildtime         time.Time `json:"buildTime"`
	Isdebug           bool      `json:"isDebug"`
	Isproduction      bool      `json:"isProduction"`
	Isadmin           bool      `json:"isAdmin"`
	Isuserinteractive bool      `json:"isUserInteractive"`
	Startuppath       string    `json:"startupPath"`
	Appdata           string    `json:"appData"`
	Osname            string    `json:"osName"`
	Osversion         string    `json:"osVersion"`
	Ismonoruntime     bool      `json:"isMonoRuntime"`
	Ismono            bool      `json:"isMono"`
	Islinux           bool      `json:"isLinux"`
	Isosx             bool      `json:"isOsx"`
	Iswindows         bool      `json:"isWindows"`
	Branch            string    `json:"branch"`
	Authentication    string    `json:"authentication"`
	Sqliteversion     string    `json:"sqliteVersion"`
	Urlbase           string    `json:"urlBase"`
	Runtimeversion    string    `json:"runtimeVersion"`
	Runtimename       string    `json:"runtimeName"`
}

type Series []struct {
	Title             string            `json:"title"`
	Alternatetitles   []Alternatetitles `json:"alternateTitles"`
	Sorttitle         string            `json:"sortTitle"`
	Seasoncount       int               `json:"seasonCount"`
	Totalepisodecount int               `json:"totalEpisodeCount,omitempty"`
	Episodecount      int               `json:"episodeCount,omitempty"`
	Episodefilecount  int               `json:"episodeFileCount,omitempty"`
	Sizeondisk        int               `json:"sizeOnDisk,omitempty"`
	Status            string            `json:"status"`
	Overview          string            `json:"overview"`
	Previousairing    time.Time         `json:"previousAiring,omitempty"`
	Network           string            `json:"network,omitempty"`
	Airtime           string            `json:"airTime,omitempty"`
	Images            []Images          `json:"images"`
	Seasons           []Seasons         `json:"seasons"`
	Year              int               `json:"year"`
	Path              string            `json:"path"`
	Profileid         int               `json:"profileId"`
	Languageprofileid int               `json:"languageProfileId"`
	Seasonfolder      bool              `json:"seasonFolder"`
	Monitored         bool              `json:"monitored"`
	Usescenenumbering bool              `json:"useSceneNumbering"`
	Runtime           int               `json:"runtime"`
	Tvdbid            int               `json:"tvdbId"`
	Tvrageid          int               `json:"tvRageId"`
	Tvmazeid          int               `json:"tvMazeId"`
	Firstaired        time.Time         `json:"firstAired,omitempty"`
	Lastinfosync      time.Time         `json:"lastInfoSync"`
	Seriestype        string            `json:"seriesType"`
	Cleantitle        string            `json:"cleanTitle"`
	Imdbid            string            `json:"imdbId,omitempty"`
	Titleslug         string            `json:"titleSlug"`
	Certification     string            `json:"certification,omitempty"`
	Genres            []string          `json:"genres"`
	Tags              []interface{}     `json:"tags"`
	Added             time.Time         `json:"added"`
	Ratings           Ratings           `json:"ratings"`
	Qualityprofileid  int               `json:"qualityProfileId"`
	ID                int               `json:"id"`
	Nextairing        time.Time         `json:"nextAiring,omitempty"`
}
type Alternatetitles struct {
	Title        string `json:"title"`
	Seasonnumber int    `json:"seasonNumber"`
}
type Images struct {
	Covertype string `json:"coverType"`
	URL       string `json:"url"`
	Remoteurl string `json:"remoteUrl"`
}

type Statistics struct {
	Previousairing    time.Time `json:"previousAiring,omitempty"`
	Episodefilecount  int       `json:"episodeFileCount"`
	Episodecount      int       `json:"episodeCount"`
	Totalepisodecount int       `json:"totalEpisodeCount"`
	Sizeondisk        int64     `json:"sizeOnDisk"`
	Percentofepisodes float64   `json:"percentOfEpisodes"`
}
type Seasons struct {
	Seasonnumber int        `json:"seasonNumber"`
	Monitored    bool       `json:"monitored"`
	Statistics   Statistics `json:"statistics,omitempty"`
}
type Ratings struct {
	Votes int     `json:"votes"`
	Value float64 `json:"value"`
}

type EpisodeFile []struct {
	Seriesid     int       `json:"seriesId"`
	Seasonnumber int       `json:"seasonNumber"`
	Path         string    `json:"path"`
	Size         int64     `json:"size"`
	Dateadded    time.Time `json:"dateAdded"`
	Scenename    string    `json:"sceneName"` //Kind of useless
	Quality      Quality   `json:"quality"`
	ID           int       `json:"id"`
}
type SubQuality struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
type Quality struct {
	Quality SubQuality `json:"quality"`
	Proper  bool       `json:"proper"`
}

type Episode struct {
	Seriesid              int       `json:"seriesId"`
	Episodefileid         int       `json:"episodeFileId"`
	Seasonnumber          int       `json:"seasonNumber"`
	Episodenumber         int       `json:"episodeNumber"`
	Title                 string    `json:"title"`
	Airdate               string    `json:"airDate"`
	Airdateutc            time.Time `json:"airDateUtc"`
	Overview              string    `json:"overview"`
	Hasfile               bool      `json:"hasFile"`
	Monitored             bool      `json:"monitored"`
	Sceneepisodenumber    int       `json:"sceneEpisodeNumber"`
	Sceneseasonnumber     int       `json:"sceneSeasonNumber"`
	Tvdbepisodeid         int       `json:"tvDbEpisodeId"`
	Absoluteepisodenumber int       `json:"absoluteEpisodeNumber"`
	ID                    int       `json:"id"`
}
