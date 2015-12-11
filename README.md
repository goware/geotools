GEOy
====

GEOy is a library and Web service written in Go that creates a singular interface to geo-location
data via the Google Maps/Location APIs.

## Part 1

Geo data comes in many forms:
* Strings, "Toronto" (a city), "Eiffel Tower" (a place)
* Geo-point, [10,20] (lng,lat)

In order to query or convert between structures, the data must be in a normalized form.

Geoy should transform between:

1. Point (lat,lng) -> Place, City or Address

2. Place, City, or Address -> Point (lat, lng)

..and I wonder where point + radius fits in?


### Social data

Pressly also queries a lot of social data from specifically, Twitter, Instagram and Facebook. Each
of these networks return a different geo-type or locality (radius).

Check out the Twitter and Instagram APIs with their sample consoles, and lets make sure for a
tweet or a post, we can map the geo-data that they return to the same form we need so we can
find all the tweets that match "food" in "Liberty Village" (a place, which is a neighbourhood
in Toronto).


## Part 2

Pressly hubs, posts, and users can all be geo-spaced.

For example:
  * User, "Peter" is from "Toronto"
  * User, "Peter" was in "Liberty Village" on Thursday Dec 10th at 9am
  * Hub, "TechCommunity" is set a point <10,20> (random lat/lng) in "Liberty Village"
  * Post, "How we all fell in love with Go" is a post at point <15,20> made in the "TechCommunity" hub

.. all of those are optional, where a hub can have a beacon point with radius 5km(?), or perhaps city / place bound,
and the posts made to the hub don't require a geo-point. But, the post would show up if a user in the hub's
beacon was exploring the area, or searching for a matching tag in the post.

The point of a hub is to anchor a community and allow the posts or users to be discovered that would like it,
but still create connections among the things around us.

```
 --------            --------            -------
|  Hub   |---------*| Post   |*---------| User  |
 --------            --------            -------
```
