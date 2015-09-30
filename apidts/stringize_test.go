package apidts

import (
	"os"
	"os/exec"
	"strings"
	"testing"
)

func convertThenStringize(s string) (string, error) {
	reader := strings.NewReader(s)
	var e error

	d, e := ConvertJsonToDts(reader)
	if e != nil {
		return "", e
	}

	r, e := StringizeDts(d, "")
	if e != nil {
		return "", e
	}

	return r, nil
}

func testCompileWithTsc(s string) error {
	code, e1 := convertThenStringize(s)
	if e1 != nil {
		return e1
	}

	t, e2 := os.Create("test.ts")
	if e2 != nil {
		panic(e2.Error())
	}
	defer t.Close()
	defer os.Remove(t.Name())

	t.WriteString(code)

	_, e3 := exec.Command("tsc", "--noEmit", t.Name()).Output()
	if e3 != nil {
		return e3
	}

	return nil
}

func TestStringizeDts(t *testing.T) {

	check := func(json string, expected string) {
		actual, e := convertThenStringize(json)
		if e != nil {
			t.Errorf("'%s' must not occur error", json)
		}
		if actual != expected {
			t.Errorf("Expected '%v' but actually '%v'", expected, actual)
		}
	}

	check(`{"foo": true}`, `interface FixMe {
  foo: boolean;
}`)
	check(`[{"foo": true}, {"foo": false}]`, `interface FixMe {
  foo: boolean;
}`)
	check(`{"bar": 42}`, `interface FixMe {
  bar: number;
}`)
	check(`{"foo": [1, 2, 3]}`, `interface FixMe {
  foo: number[];
}`)
	check(`{"foo": {"poyo": [1, 2, 3]}}`, `interface FixMe {
  foo: {
    poyo: number[];
  };
}`)
	check(`{"foo": {"poyo": {"puyo": [true]}}}`, `interface FixMe {
  foo: {
    poyo: {
      puyo: boolean[];
    };
  };
}`)
	check(`{"foo": []}`, `interface FixMe {
  foo: any[];
}`)
	check(`{"foo": null}`, `interface FixMe {
  foo: any;
}`)

	if _, e := convertThenStringize("true"); e == nil {
		t.Errorf("Can't convert 'true' to interface but error does't occur")
	}

	if _, e := convertThenStringize("42"); e == nil {
		t.Errorf("Can't convert 'true' to interface but error does't occur")
	}

	if _, e := convertThenStringize("null"); e == nil {
		t.Errorf("Can't convert 'true' to interface but error does't occur")
	}

	testCompileWithTsc(`{"foo": "aaa", "bar": {"poyo": [true, false, true], "puyo": {"aaa": 42}}}`)
	testCompileWithTsc(`[
{
  "contributors": null,
  "coordinates": null,
  "created_at": "Tue Sep 01 08:40:28 +0000 2015",
  "entities": {
    "hashtags": [],
    "symbols": [],
    "urls": [
      {
        "display_url": "twitter.com/Linda_pp/statu…",
        "expanded_url": "https://twitter.com/Linda_pp/status/638402862934986752",
        "indices": [
          6,
          29
        ],
        "url": "https://t.co/vZil2mnSvQ"
      }
    ],
    "user_mentions": []
  },
  "favorite_count": 0,
  "favorited": false,
  "filter_level": "low",
  "geo": null,
  "id": 638632504522444800,
  "id_str": "638632504522444800",
  "in_reply_to_screen_name": null,
  "in_reply_to_status_id": null,
  "in_reply_to_status_id_str": null,
  "in_reply_to_user_id": null,
  "in_reply_to_user_id_str": null,
  "is_quote_status": true,
  "lang": "ja",
  "place": null,
  "possibly_sensitive": false,
  "quoted_status": {
    "contributors": null,
    "coordinates": null,
    "created_at": "Mon Aug 31 17:27:58 +0000 2015",
    "entities": {
      "hashtags": [],
      "symbols": [],
      "urls": [],
      "user_mentions": []
    },
    "favorite_count": 0,
    "favorited": false,
    "filter_level": "low",
    "geo": null,
    "id": 638402862934986800,
    "id_str": "638402862934986752",
    "in_reply_to_screen_name": null,
    "in_reply_to_status_id": null,
    "in_reply_to_status_id_str": null,
    "in_reply_to_user_id": null,
    "in_reply_to_user_id_str": null,
    "is_quote_status": false,
    "lang": "ja",
    "place": null,
    "retweet_count": 0,
    "retweeted": false,
    "source": "<a href=\"http://sites.google.com/site/yorufukurou/\" rel=\"nofollow\">YoruFukurou</a>",
    "text": "眠すぎるため寝ます",
    "truncated": false,
    "user": {
      "contributors_enabled": false,
      "created_at": "Thu Mar 04 17:10:18 +0000 2010",
      "default_profile": false,
      "default_profile_image": false,
      "description": "ソフトウェアエンジニア見習い．趣味で C++ (C++11 or later)，Ruby，Dachs をVimったりする．計算機言語などのプログラミングツールが好き．Electron + TypeScript でデスクトップアプリ始めました．あと写真も楽しい．犬．",
      "favourites_count": 384,
      "follow_request_sent": null,
      "followers_count": 1297,
      "following": null,
      "friends_count": 373,
      "geo_enabled": false,
      "id": 119789510,
      "id_str": "119789510",
      "is_translator": false,
      "lang": "en",
      "listed_count": 149,
      "location": "Tokyo ^ Kanagawa",
      "name": "ドッグ",
      "notifications": null,
      "profile_background_color": "B3B3B3",
      "profile_background_image_url": "http://pbs.twimg.com/profile_background_images/458967069522817025/VbYAPpF5.png",
      "profile_background_image_url_https": "https://pbs.twimg.com/profile_background_images/458967069522817025/VbYAPpF5.png",
      "profile_background_tile": true,
      "profile_banner_url": "https://pbs.twimg.com/profile_banners/119789510/1367930390",
      "profile_image_url": "http://pbs.twimg.com/profile_images/3626384430/3a64cf406665c1940d68ab737003605c_normal.jpeg",
      "profile_image_url_https": "https://pbs.twimg.com/profile_images/3626384430/3a64cf406665c1940d68ab737003605c_normal.jpeg",
      "profile_link_color": "545454",
      "profile_sidebar_border_color": "FFFFFF",
      "profile_sidebar_fill_color": "E6E6E6",
      "profile_text_color": "050505",
      "profile_use_background_image": true,
      "protected": false,
      "screen_name": "Linda_pp",
      "statuses_count": 126429,
      "time_zone": "Osaka",
      "url": "https://github.com/rhysd",
      "utc_offset": 32400,
      "verified": false
    }
  },
  "quoted_status_id": 638402862934986800,
  "quoted_status_id_str": "638402862934986752",
  "retweet_count": 0,
  "retweeted": false,
  "source": "<a href=\"https://twitter.com/#!/Linda_pp\" rel=\"nofollow\">犬Vim</a>",
  "text": "テストです https://t.co/vZil2mnSvQ",
  "timestamp_ms": "1441096828953",
  "truncated": false,
  "user": {
    "contributors_enabled": false,
    "created_at": "Thu Mar 04 17:10:18 +0000 2010",
    "default_profile": false,
    "default_profile_image": false,
    "description": "ソフトウェアエンジニア見習い．趣味で C++ (C++11 or later)，Ruby，Dachs をVimったりする．計算機言語などのプログラミングツールが好き．Electron + TypeScript でデスクトップアプリ始めました．あと写真も楽しい．犬．",
    "favourites_count": 384,
    "follow_request_sent": null,
    "followers_count": 1297,
    "following": null,
    "friends_count": 373,
    "geo_enabled": false,
    "id": 119789510,
    "id_str": "119789510",
    "is_translator": false,
    "lang": "en",
    "listed_count": 149,
    "location": "Tokyo ^ Kanagawa",
    "name": "ドッグ",
    "notifications": null,
    "profile_background_color": "B3B3B3",
    "profile_background_image_url": "http://pbs.twimg.com/profile_background_images/458967069522817025/VbYAPpF5.png",
    "profile_background_image_url_https": "https://pbs.twimg.com/profile_background_images/458967069522817025/VbYAPpF5.png",
    "profile_background_tile": true,
    "profile_banner_url": "https://pbs.twimg.com/profile_banners/119789510/1367930390",
    "profile_image_url": "http://pbs.twimg.com/profile_images/3626384430/3a64cf406665c1940d68ab737003605c_normal.jpeg",
    "profile_image_url_https": "https://pbs.twimg.com/profile_images/3626384430/3a64cf406665c1940d68ab737003605c_normal.jpeg",
    "profile_link_color": "545454",
    "profile_sidebar_border_color": "FFFFFF",
    "profile_sidebar_fill_color": "E6E6E6",
    "profile_text_color": "050505",
    "profile_use_background_image": true,
    "protected": false,
    "screen_name": "Linda_pp",
    "statuses_count": 126430,
    "time_zone": "Osaka",
    "url": "https://github.com/rhysd",
    "utc_offset": 32400,
    "verified": false
  }
}
]`)
}
