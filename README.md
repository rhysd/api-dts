`d.ts` generator using JSON API response
========================================
[![Build Status](https://travis-ci.org/rhysd/api-dts.svg)](https://travis-ci.org/rhysd/api-dts)

`api-dts` is a generator for TypeScript programmer who use some JSON APIs.  API response JSON has too many fields to write the type definition for it manually.  `api-dts` generates such an annoying type definition automatically.

```
$ api-dts some-api.json > some-api.d.ts
```

`api-dts` reads JSON text from file specified as argument and simply writes the result to STDOUT.  If command argument is ommited, `api-dts` reads STDIN.  `api-dts` defines interface name of the API from the specified file name, so specifying `-out` prefers to redirecting to file.
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

`$ api-dts my-api.json > my-api.d.ts` generates below type definition.

```typescript
interface MyApi {
  user: {
    age: number;
    lang: string;
    name: string;
  };
  progress: boolean;
}
```

## TODO

- Seprate sub interfaces.  Their names are made using the key names of them.
- Detect optional field (suffix `?`)
- When the JSON is an array, check all elements have the same interface

## Real World Example

In general, API document shows an example response.  You can simply copy it.

For example, below is `GET users/show` Twitter API response shown in [document](https://dev.twitter.com/rest/reference/get/users/show).  Assume that you copied it as `twitter-user.json`.

```json
{
  "contributors_enabled": false,
  "created_at": "Sat Dec 14 04:35:55 +0000 2013",
  "default_profile": false,
  "default_profile_image": false,
  "description": "Developer and Platform Relations @Twitter. We are developer advocates. We can't answer all your questions, but we listen to all of them!",
  "entities": {
    "description": {
      "urls": []
    },
    "url": {
      "urls": [
        {
          "display_url": "dev.twitter.com",
          "expanded_url": "https://dev.twitter.com/",
          "indices": [
            0,
            23
          ],
          "url": "https://t.co/66w26cua1O"
        }
      ]
    }
  },
  "favourites_count": 757,
  "follow_request_sent": false,
  "followers_count": 143916,
  "following": false,
  "friends_count": 1484,
  "geo_enabled": true,
  "id": 2244994945,
  "id_str": "2244994945",
  "is_translation_enabled": false,
  "is_translator": false,
  "lang": "en",
  "listed_count": 516,
  "location": "Internet",
  "name": "TwitterDev",
  "notifications": false,
  "profile_background_color": "FFFFFF",
  "profile_background_image_url": "http://abs.twimg.com/images/themes/theme1/bg.png",
  "profile_background_image_url_https": "https://abs.twimg.com/images/themes/theme1/bg.png",
  "profile_background_tile": false,
  "profile_banner_url": "https://pbs.twimg.com/profile_banners/2244994945/1396995246",
  "profile_image_url": "http://pbs.twimg.com/profile_images/530814764687949824/npQQVkq8_normal.png",
  "profile_image_url_https": "https://pbs.twimg.com/profile_images/530814764687949824/npQQVkq8_normal.png",
  "profile_link_color": "0084B4",
  "profile_location": null,
  "profile_sidebar_border_color": "FFFFFF",
  "profile_sidebar_fill_color": "DDEEF6",
  "profile_text_color": "333333",
  "profile_use_background_image": false,
  "protected": false,
  "screen_name": "TwitterDev",
  "status": {
    "contributors": null,
    "coordinates": null,
    "created_at": "Fri Jun 12 19:50:18 +0000 2015",
    "entities": {
      "hashtags": [],
      "symbols": [],
      "urls": [
        {
          "display_url": "github.com/twitterdev/twi\u2026",
          "expanded_url": "https://github.com/twitterdev/twitter-for-bigquery",
          "indices": [
            36,
            59
          ],
          "url": "https://t.co/K5orgXzhOM"
        }
      ],
      "user_mentions": [
        {
          "id": 18518601,
          "id_str": "18518601",
          "indices": [
            3,
            13
          ],
          "name": "William Vambenepe",
          "screen_name": "vambenepe"
        }
      ]
    },
    "favorite_count": 0,
    "favorited": false,
    "geo": null,
    "id": 609447655429787648,
    "id_str": "609447655429787648",
    "in_reply_to_screen_name": null,
    "in_reply_to_status_id": null,
    "in_reply_to_status_id_str": null,
    "in_reply_to_user_id": null,
    "in_reply_to_user_id_str": null,
    "lang": "en",
    "place": null,
    "possibly_sensitive": false,
    "retweet_count": 19,
    "retweeted": false,
    "retweeted_status": {
      "contributors": null,
      "coordinates": null,
      "created_at": "Fri Jun 12 05:19:11 +0000 2015",
      "entities": {
        "hashtags": [],
        "symbols": [],
        "urls": [
          {
            "display_url": "github.com/twitterdev/twi\u2026",
            "expanded_url": "https://github.com/twitterdev/twitter-for-bigquery",
            "indices": [
              21,
              44
            ],
            "url": "https://t.co/K5orgXzhOM"
          }
        ],
        "user_mentions": []
      },
      "favorite_count": 23,
      "favorited": false,
      "geo": null,
      "id": 609228428915552257,
      "id_str": "609228428915552257",
      "in_reply_to_screen_name": null,
      "in_reply_to_status_id": null,
      "in_reply_to_status_id_str": null,
      "in_reply_to_user_id": null,
      "in_reply_to_user_id_str": null,
      "lang": "en",
      "place": null,
      "possibly_sensitive": false,
      "retweet_count": 19,
      "retweeted": false,
      "source": "<a>Twitter Web Client</a>",
      "text": "Twitter for BigQuery https://t.co/K5orgXzhOM See how easy it is to stream Twitter data into BigQuery.",
      "truncated": false
    },
    "source": "<a>Twitter for iPhone</a>",
    "text": "RT @vambenepe: Twitter for BigQuery https://t.co/K5orgXzhOM See how easy it is to stream Twitter data into BigQuery.",
    "truncated": false
  },
  "statuses_count": 1279,
  "time_zone": "Pacific Time (US & Canada)",
  "url": "https://t.co/66w26cua1O",
  "utc_offset": -25200,
  "verified": true
}
```

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
$ api-dts twitter-user.json > twitter-user.d.ts
```

It writes type definition to `twitter-user.d.ts` for the API as below.

```typescript
interface TwitterUser {
  time_zone: string;
  created_at: string;
  screen_name: string;
  following: boolean;
  listed_count: number;
  description: string;
  id: number;
  profile_background_color: string;
  location: string;
  default_profile: boolean;
  is_translator: boolean;
  profile_background_image_url_https: string;
  statuses_count: number;
  name: string;
  profile_text_color: string;
  contributors_enabled: boolean;
  profile_banner_url: string;
  profile_image_url_https: string;
  friends_count: number;
  profile_link_color: string;
  geo_enabled: boolean;
  is_translation_enabled: boolean;
  favourites_count: number;
  notifications: boolean;
  profile_background_tile: boolean;
  profile_image_url: string;
  utc_offset: number;
  profile_sidebar_fill_color: string;
  protected: boolean;
  profile_location: any;
  lang: string;
  default_profile_image: boolean;
  id_str: string;
  status: {
    contributors: any;
    id: number;
    in_reply_to_user_id: any;
    retweet_count: number;
    truncated: boolean;
    possibly_sensitive: boolean;
    source: string;
    geo: any;
    place: any;
    retweeted_status: {
      favorite_count: number;
      geo: any;
      in_reply_to_status_id: any;
      possibly_sensitive: boolean;
      truncated: boolean;
      created_at: string;
      text: string;
      entities: {
        user_mentions: any[];
        hashtags: any[];
        symbols: any[];
        urls: {
          expanded_url: string;
          indices: number[];
          url: string;
          display_url: string;
        }[];
      };
      favorited: boolean;
      in_reply_to_user_id: any;
      retweeted: boolean;
      in_reply_to_user_id_str: any;
      contributors: any;
      coordinates: any;
      place: any;
      retweet_count: number;
      source: string;
      in_reply_to_status_id_str: any;
      lang: string;
      id_str: string;
      in_reply_to_screen_name: any;
      id: number;
    };
    text: string;
    retweeted: boolean;
    created_at: string;
    in_reply_to_status_id: any;
    lang: string;
    coordinates: any;
    favorite_count: number;
    entities: {
      urls: {
        display_url: string;
        expanded_url: string;
        indices: number[];
        url: string;
      }[];
      user_mentions: {
        id: number;
        id_str: string;
        indices: number[];
        name: string;
        screen_name: string;
      }[];
      hashtags: any[];
      symbols: any[];
    };
    in_reply_to_screen_name: any;
    id_str: string;
    in_reply_to_status_id_str: any;
    favorited: boolean;
    in_reply_to_user_id_str: any;
  };
  profile_sidebar_border_color: string;
  profile_background_image_url: string;
  url: string;
  entities: {
    url: {
      urls: {
        expanded_url: string;
        indices: number[];
        url: string;
        display_url: string;
      }[];
    };
    description: {
      urls: any[];
    };
  };
  followers_count: number;
  profile_use_background_image: boolean;
  follow_request_sent: boolean;
  verified: boolean;
}
```
