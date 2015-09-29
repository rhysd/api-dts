`d.ts` generator using JSON API response
========================================
[![Build Status](https://travis-ci.org/rhysd/api-dts.svg)](https://travis-ci.org/rhysd/api-dts)

`api-dts` is a generator for TypeScript programmer who use some JSON APIs.  API response JSON has too many fields to write the type definition for it manually.  `api-dts` generates such an annoying type definition automatically.

```
$ api-dts < response.json > api.d.ts
```

`api-dts` simply reads STDIN and writes result to STDOUT.
You can install `api-dts` with `go get`.

```
go get github.com/rhysd/api-dts
```

## Example

Assume that below JSON is API response.

```json
[
  {
    "user": {
      "name": "rhysd",
      "age": 27,
      "lang": "Dachs"
    },
    "has_progress": false
  },
  {
    "user": {
      "name": "linda",
      "age": 24,
      "lang": "scala"
    },
    "has_progress": true
  }
]
```

`api-dts` generates below type definition.

```typescript
interface FixMe  {
  user: {
    age: number;
    lang: string;
    name: string;
  };
  progress: boolean;
}
```

You can save it to `{api}.d.ts` and rename `FixMe` to the name of API.

## TODO

- Add `-o`(`--out`) option and make name of the interface from it.
- Seprate sub interfaces.  Their names are made using the key names of them.
- Detect optional field (suffix `?`)
- When the JSON is an array, check all elements have the same interface

## Real World Example

Below is a response from Twitter API

```json
[
  {
    "contributors": null,
    "coordinates": null,
    "created_at": "Wed Sep 02 02:11:45 +0000 2015",
    "entities": {
      "hashtags": [],
      "media": [
        {
          "display_url": "pic.twitter.com/bb5t5iNjiL",
          "expanded_url": "http://twitter.com/msrn_Olll/status/638566389972799488/photo/1",
          "id": 638566382586585100,
          "id_str": "638566382586585088",
          "indices": [
            33,
            55
          ],
          "media_url": "http://pbs.twimg.com/media/CNykvd-UYAAzURE.jpg",
          "media_url_https": "https://pbs.twimg.com/media/CNykvd-UYAAzURE.jpg",
          "sizes": {
            "large": {
              "h": 576,
              "resize": "fit",
              "w": 1024
            },
            "medium": {
              "h": 337,
              "resize": "fit",
              "w": 600
            },
            "small": {
              "h": 191,
              "resize": "fit",
              "w": 340
            },
            "thumb": {
              "h": 150,
              "resize": "crop",
              "w": 150
            }
          },
          "source_status_id": 638566389972799500,
          "source_status_id_str": "638566389972799488",
          "source_user_id": 2931268554,
          "source_user_id_str": "2931268554",
          "type": "photo",
          "url": "http://t.co/bb5t5iNjiL"
        }
      ],
      "symbols": [],
      "urls": [],
      "user_mentions": [
        {
          "id": 2931268554,
          "id_str": "2931268554",
          "indices": [
            3,
            13
          ],
          "name": "まーくん",
          "screen_name": "msrn_Olll"
        }
      ]
    },
    "extended_entities": {
      "media": [
        {
          "display_url": "pic.twitter.com/bb5t5iNjiL",
          "expanded_url": "http://twitter.com/msrn_Olll/status/638566389972799488/photo/1",
          "id": 638566382586585100,
          "id_str": "638566382586585088",
          "indices": [
            33,
            55
          ],
          "media_url": "http://pbs.twimg.com/media/CNykvd-UYAAzURE.jpg",
          "media_url_https": "https://pbs.twimg.com/media/CNykvd-UYAAzURE.jpg",
          "sizes": {
            "large": {
              "h": 576,
              "resize": "fit",
              "w": 1024
            },
            "medium": {
              "h": 337,
              "resize": "fit",
              "w": 600
            },
            "small": {
              "h": 191,
              "resize": "fit",
              "w": 340
            },
            "thumb": {
              "h": 150,
              "resize": "crop",
              "w": 150
            }
          },
          "source_status_id": 638566389972799500,
          "source_status_id_str": "638566389972799488",
          "source_user_id": 2931268554,
          "source_user_id_str": "2931268554",
          "type": "photo",
          "url": "http://t.co/bb5t5iNjiL"
        },
        {
          "display_url": "pic.twitter.com/bb5t5iNjiL",
          "expanded_url": "http://twitter.com/msrn_Olll/status/638566389972799488/photo/1",
          "id": 638566388131459100,
          "id_str": "638566388131459073",
          "indices": [
            33,
            55
          ],
          "media_url": "http://pbs.twimg.com/media/CNykvyoUcAEmyhb.jpg",
          "media_url_https": "https://pbs.twimg.com/media/CNykvyoUcAEmyhb.jpg",
          "sizes": {
            "large": {
              "h": 576,
              "resize": "fit",
              "w": 1024
            },
            "medium": {
              "h": 337,
              "resize": "fit",
              "w": 600
            },
            "small": {
              "h": 191,
              "resize": "fit",
              "w": 340
            },
            "thumb": {
              "h": 150,
              "resize": "crop",
              "w": 150
            }
          },
          "source_status_id": 638566389972799500,
          "source_status_id_str": "638566389972799488",
          "source_user_id": 2931268554,
          "source_user_id_str": "2931268554",
          "type": "photo",
          "url": "http://t.co/bb5t5iNjiL"
        }
      ]
    },
    "favorite_count": 0,
    "favorited": false,
    "filter_level": "low",
    "geo": null,
    "id": 638897065146224600,
    "id_str": "638897065146224640",
    "in_reply_to_screen_name": null,
    "in_reply_to_status_id": null,
    "in_reply_to_status_id_str": null,
    "in_reply_to_user_id": null,
    "in_reply_to_user_id_str": null,
    "is_quote_status": false,
    "lang": "ja",
    "place": null,
    "possibly_sensitive": false,
    "retweet_count": 0,
    "retweeted": false,
    "retweeted_status": {
      "contributors": null,
      "coordinates": null,
      "created_at": "Tue Sep 01 04:17:46 +0000 2015",
      "entities": {
        "hashtags": [],
        "media": [
          {
            "display_url": "pic.twitter.com/bb5t5iNjiL",
            "expanded_url": "http://twitter.com/msrn_Olll/status/638566389972799488/photo/1",
            "id": 638566382586585100,
            "id_str": "638566382586585088",
            "indices": [
              18,
              40
            ],
            "media_url": "http://pbs.twimg.com/media/CNykvd-UYAAzURE.jpg",
            "media_url_https": "https://pbs.twimg.com/media/CNykvd-UYAAzURE.jpg",
            "sizes": {
              "large": {
                "h": 576,
                "resize": "fit",
                "w": 1024
              },
              "medium": {
                "h": 337,
                "resize": "fit",
                "w": 600
              },
              "small": {
                "h": 191,
                "resize": "fit",
                "w": 340
              },
              "thumb": {
                "h": 150,
                "resize": "crop",
                "w": 150
              }
            },
            "type": "photo",
            "url": "http://t.co/bb5t5iNjiL"
          }
        ],
        "symbols": [],
        "urls": [],
        "user_mentions": []
      },
      "extended_entities": {
        "media": [
          {
            "display_url": "pic.twitter.com/bb5t5iNjiL",
            "expanded_url": "http://twitter.com/msrn_Olll/status/638566389972799488/photo/1",
            "id": 638566382586585100,
            "id_str": "638566382586585088",
            "indices": [
              18,
              40
            ],
            "media_url": "http://pbs.twimg.com/media/CNykvd-UYAAzURE.jpg",
            "media_url_https": "https://pbs.twimg.com/media/CNykvd-UYAAzURE.jpg",
            "sizes": {
              "large": {
                "h": 576,
                "resize": "fit",
                "w": 1024
              },
              "medium": {
                "h": 337,
                "resize": "fit",
                "w": 600
              },
              "small": {
                "h": 191,
                "resize": "fit",
                "w": 340
              },
              "thumb": {
                "h": 150,
                "resize": "crop",
                "w": 150
              }
            },
            "type": "photo",
            "url": "http://t.co/bb5t5iNjiL"
          },
          {
            "display_url": "pic.twitter.com/bb5t5iNjiL",
            "expanded_url": "http://twitter.com/msrn_Olll/status/638566389972799488/photo/1",
            "id": 638566388131459100,
            "id_str": "638566388131459073",
            "indices": [
              18,
              40
            ],
            "media_url": "http://pbs.twimg.com/media/CNykvyoUcAEmyhb.jpg",
            "media_url_https": "https://pbs.twimg.com/media/CNykvyoUcAEmyhb.jpg",
            "sizes": {
              "large": {
                "h": 576,
                "resize": "fit",
                "w": 1024
              },
              "medium": {
                "h": 337,
                "resize": "fit",
                "w": 600
              },
              "small": {
                "h": 191,
                "resize": "fit",
                "w": 340
              },
              "thumb": {
                "h": 150,
                "resize": "crop",
                "w": 150
              }
            },
            "type": "photo",
            "url": "http://t.co/bb5t5iNjiL"
          }
        ]
      },
      "favorite_count": 821,
      "favorited": true,
      "filter_level": "low",
      "geo": null,
      "id": 638566389972799500,
      "id_str": "638566389972799488",
      "in_reply_to_screen_name": null,
      "in_reply_to_status_id": null,
      "in_reply_to_status_id_str": null,
      "in_reply_to_user_id": null,
      "in_reply_to_user_id_str": null,
      "is_quote_status": false,
      "lang": "ja",
      "place": null,
      "possibly_sensitive": false,
      "retweet_count": 693,
      "retweeted": true,
      "source": "<a href=\"http://twitter.com/download/android\" rel=\"nofollow\">Twitter for Android</a>",
      "text": "ε   -   δ   論   法 http://t.co/bb5t5iNjiL",
      "truncated": false,
      "user": {
        "contributors_enabled": false,
        "created_at": "Mon Dec 15 14:53:16 +0000 2014",
        "default_profile": true,
        "default_profile_image": false,
        "description": "大阪桐蔭にいたよ。京大にいるよ。",
        "favourites_count": 1895,
        "follow_request_sent": null,
        "followers_count": 118,
        "following": null,
        "friends_count": 130,
        "geo_enabled": true,
        "id": 2931268554,
        "id_str": "2931268554",
        "is_translator": false,
        "lang": "ja",
        "listed_count": 2,
        "location": "イデヤ界",
        "name": "まーくん",
        "notifications": null,
        "profile_background_color": "C0DEED",
        "profile_background_image_url": "http://abs.twimg.com/images/themes/theme1/bg.png",
        "profile_background_image_url_https": "https://abs.twimg.com/images/themes/theme1/bg.png",
        "profile_background_tile": false,
        "profile_banner_url": "https://pbs.twimg.com/profile_banners/2931268554/1440482341",
        "profile_image_url": "http://pbs.twimg.com/profile_images/634729243977498625/ardXs_EI_normal.jpg",
        "profile_image_url_https": "https://pbs.twimg.com/profile_images/634729243977498625/ardXs_EI_normal.jpg",
        "profile_link_color": "0084B4",
        "profile_sidebar_border_color": "C0DEED",
        "profile_sidebar_fill_color": "DDEEF6",
        "profile_text_color": "333333",
        "profile_use_background_image": true,
        "protected": false,
        "screen_name": "msrn_Olll",
        "statuses_count": 13830,
        "time_zone": null,
        "url": null,
        "utc_offset": null,
        "verified": false
      }
    },
    "source": "<a href=\"http://twitter.com/download/iphone\" rel=\"nofollow\">Twitter for iPhone</a>",
    "text": "RT @msrn_Olll: ε   -   δ   論   法 http://t.co/bb5t5iNjiL",
    "timestamp_ms": "1441159905121",
    "truncated": false,
    "user": {
      "contributors_enabled": false,
      "created_at": "Thu Mar 07 10:31:07 +0000 2013",
      "default_profile": true,
      "default_profile_image": false,
      "description": "一つ、考え事をしよう。 市大工電情、ハミデント。Javaの勉強始めました。たまに呟くのでアドバイスとかくれるとありがたいです。ADOCUS立ち上げプロジェクト進行中。感傷ベクトルとかibとか入間人間とか。",
      "favourites_count": 4844,
      "follow_request_sent": null,
      "followers_count": 107,
      "following": null,
      "friends_count": 130,
      "geo_enabled": false,
      "id": 1248415158,
      "id_str": "1248415158",
      "is_translator": false,
      "lang": "ja",
      "listed_count": 3,
      "location": "新中が見えそうで見えない。",
      "name": "KK",
      "notifications": null,
      "profile_background_color": "C0DEED",
      "profile_background_image_url": "http://abs.twimg.com/images/themes/theme1/bg.png",
      "profile_background_image_url_https": "https://abs.twimg.com/images/themes/theme1/bg.png",
      "profile_background_tile": false,
      "profile_banner_url": "https://pbs.twimg.com/profile_banners/1248415158/1435703212",
      "profile_image_url": "http://pbs.twimg.com/profile_images/592143917841125376/TtfvA37y_normal.jpg",
      "profile_image_url_https": "https://pbs.twimg.com/profile_images/592143917841125376/TtfvA37y_normal.jpg",
      "profile_link_color": "0084B4",
      "profile_sidebar_border_color": "C0DEED",
      "profile_sidebar_fill_color": "DDEEF6",
      "profile_text_color": "333333",
      "profile_use_background_image": true,
      "protected": false,
      "screen_name": "AS_predri",
      "statuses_count": 5319,
      "time_zone": "Tokyo",
      "url": null,
      "utc_offset": 32400,
      "verified": false
    }
  }
]
```

Then execute

```
$ api-dts < response.json
```

It outputs type definition for the API.  You can save it to `{some-api}.d.ts`.  Left task is only to rename the name of interface.

```typescript
interface FixMe  {
  in_reply_to_status_id: any;
  extended_entities: {
    media: {
      source_user_id_str: string;
      source_status_id_str: string;
      id_str: string;
      indices: number[];
      media_url: string;
      media_url_https: string;
      source_user_id: number;
      type: string;
      url: string;
      expanded_url: string;
      source_status_id: number;
      id: number;
      sizes: {
        medium: {
          h: number;
          resize: string;
          w: number;
        };
        small: {
          h: number;
          resize: string;
          w: number;
        };
        thumb: {
          resize: string;
          w: number;
          h: number;
        };
        large: {
          h: number;
          resize: string;
          w: number;
        };
      };
      display_url: string;
    }[];
  };
  contributors: any;
  timestamp_ms: string;
  user: {
    profile_background_image_url: string;
    profile_background_image_url_https: string;
    profile_link_color: string;
    profile_sidebar_fill_color: string;
    favourites_count: number;
    screen_name: string;
    contributors_enabled: boolean;
    default_profile_image: boolean;
    id_str: string;
    listed_count: number;
    time_zone: string;
    url: any;
    default_profile: boolean;
    profile_use_background_image: boolean;
    profile_sidebar_border_color: string;
    is_translator: boolean;
    follow_request_sent: any;
    location: string;
    profile_image_url: string;
    friends_count: number;
    geo_enabled: boolean;
    profile_background_tile: boolean;
    utc_offset: number;
    profile_background_color: string;
    name: string;
    profile_image_url_https: string;
    statuses_count: number;
    lang: string;
    description: string;
    following: any;
    protected: boolean;
    followers_count: number;
    profile_banner_url: string;
    verified: boolean;
    profile_text_color: string;
    id: number;
    notifications: any;
    created_at: string;
  };
  is_quote_status: boolean;
  in_reply_to_screen_name: any;
  possibly_sensitive: boolean;
  in_reply_to_user_id: any;
  retweeted_status: {
    contributors: any;
    entities: {
      symbols: any[];
      urls: any[];
      user_mentions: any[];
      hashtags: any[];
      media: {
        id: number;
        id_str: string;
        indices: number[];
        type: string;
        media_url_https: string;
        media_url: string;
        url: string;
        display_url: string;
        expanded_url: string;
        sizes: {
          thumb: {
            resize: string;
            w: number;
            h: number;
          };
          large: {
            resize: string;
            w: number;
            h: number;
          };
          medium: {
            h: number;
            resize: string;
            w: number;
          };
          small: {
            resize: string;
            w: number;
            h: number;
          };
        };
      }[];
    };
    favorited: boolean;
    in_reply_to_screen_name: any;
    id: number;
    place: any;
    in_reply_to_status_id_str: any;
    retweeted: boolean;
    source: string;
    in_reply_to_user_id: any;
    in_reply_to_user_id_str: any;
    is_quote_status: boolean;
    favorite_count: number;
    extended_entities: {
      media: {
        indices: number[];
        display_url: string;
        url: string;
        expanded_url: string;
        id: number;
        media_url_https: string;
        sizes: {
          small: {
            h: number;
            resize: string;
            w: number;
          };
          thumb: {
            h: number;
            resize: string;
            w: number;
          };
          large: {
            h: number;
            resize: string;
            w: number;
          };
          medium: {
            h: number;
            resize: string;
            w: number;
          };
        };
        type: string;
        id_str: string;
        media_url: string;
      }[];
    };
    retweet_count: number;
    id_str: string;
    coordinates: any;
    created_at: string;
    possibly_sensitive: boolean;
    in_reply_to_status_id: any;
    user: {
      protected: boolean;
      profile_background_tile: boolean;
      profile_use_background_image: boolean;
      profile_banner_url: string;
      verified: boolean;
      default_profile_image: boolean;
      id_str: string;
      follow_request_sent: any;
      profile_sidebar_fill_color: string;
      is_translator: boolean;
      profile_text_color: string;
      time_zone: any;
      url: any;
      geo_enabled: boolean;
      profile_link_color: string;
      profile_background_image_url_https: string;
      profile_background_color: string;
      friends_count: number;
      favourites_count: number;
      listed_count: number;
      notifications: any;
      lang: string;
      followers_count: number;
      profile_image_url: string;
      id: number;
      default_profile: boolean;
      statuses_count: number;
      description: string;
      screen_name: string;
      utc_offset: any;
      name: string;
      profile_sidebar_border_color: string;
      created_at: string;
      following: any;
      contributors_enabled: boolean;
      location: string;
      profile_background_image_url: string;
      profile_image_url_https: string;
    };
    truncated: boolean;
    filter_level: string;
    geo: any;
    text: string;
    lang: string;
  };
  created_at: string;
  retweet_count: number;
  favorited: boolean;
  retweeted: boolean;
  coordinates: any;
  id_str: string;
  geo: any;
  id: number;
  text: string;
  truncated: boolean;
  in_reply_to_user_id_str: any;
  source: string;
  place: any;
  filter_level: string;
  in_reply_to_status_id_str: any;
  entities: {
    media: {
      expanded_url: string;
      source_status_id: number;
      source_user_id: number;
      id_str: string;
      media_url_https: string;
      display_url: string;
      url: string;
      media_url: string;
      sizes: {
        large: {
          h: number;
          resize: string;
          w: number;
        };
        medium: {
          h: number;
          resize: string;
          w: number;
        };
        small: {
          h: number;
          resize: string;
          w: number;
        };
        thumb: {
          resize: string;
          w: number;
          h: number;
        };
      };
      id: number;
      indices: number[];
      source_user_id_str: string;
      type: string;
      source_status_id_str: string;
    }[];
    symbols: any[];
    urls: any[];
    user_mentions: {
      id: number;
      id_str: string;
      indices: number[];
      name: string;
      screen_name: string;
    }[];
    hashtags: any[];
  };
  lang: string;
  favorite_count: number;
}
```
