package meetup

import (
	"fmt"
	"log"
)

type ExplodedDate struct {
	Day	int `json:"day"`
	Month int `json:"month"`
	Year int	`json:"year"`
}

type Membership struct {
	Created	int	`json:"created"`
	Group	Group	`json:"group"`
	Role	string	`json:"role"`
	Status	string	`json:"status"`
	Title	string	`json:"title"`
	Updated	int	`json:"updated"`
	Visited	int	`json:"visited"`
}

type Memberships struct {
	Member	[]*Membership	`json:"member"`
	Organizer	[]*Membership	`json:"organizer"`
}

type OtherService struct {
	Identifier	string	`json:"identifier"`
	Url	string	`json:"url"`
}

type Photo struct {
	BaseUrl	string	`json:"base_url"`
	HighresLink	string	`json:"highres_link"`
	ID		int	`json:"id"`
	PhotoLink	string	`json:"photo_link"`
	ThumbLink	string	`json:"thumb_link"`
	Type	string	`json:"type"`
}

type Privacy struct {
	Bio	string	`json:"bio"`
	Facebook	string	`json:"facebook"`
	Groups		string	`json:"groups"`
	Topics		string	`json:"topics"`
}

type Stats struct {
	Groups		int	`json:"groups"`
	Rsvps		int	`json:"rsvps"`
	Topics		int	`json:"topics"`
}

type Host struct {
	ID			int	`json:"id"`
	Intro		string	`json:"intro"`
	Name		string	`json:"name"`
	Photo		Photo	`json:"photo"`
}

type RsvpRules struct {
	CloseTime	int	`json:"closeTime"`
	Closed		bool	`json:"closed"`
	GuestLimit	int	`json:"guestLimit"`
	OpenTime	int	`json:"openTime"`
	RefundPolicy	struct {
		Days	int	`json:"days"`
		Notes	string	`json:"notes"`
		Policies	string	`json:"policies"`
	} `json:"refundPolicy"`
	WaitListing	string	`json:"waitListing"`
}

type Rsvps struct {
	ID				string	`json:"id"`
	CommentCount	int	`json:"comment_count"`
	Created			int	`json:"created"`
	Description		string	`json:"description"`
	DescriptionImages	string`json:"description_images"`
	Duration		int	`json:"duration"`
	EventHosts		[]Host	`json:"hosts"`
	Featured		bool	`json:"featured"`
	FeaturedPhoto	Photo	`json:"featured_photo"`
	Fee				EventFee	`json:"event_fee"`
	Group			Group	`json:"group"`
	HowToFindUs		string	`json:"how_to_find_us"`
	Link			string	`json:"link"`
	LocalDate		string	`json:"local_date"`
	LocalTime		string	`json:"local_time"`
	ManualAttendanceCount	int	`json:"manual_attendance_count"`
	Name			string	`json:"name"`
	// PhotoAlbum
	PlainTextDescription	string	`json:"plain_text_description"`
	PlainTextNoImagesDescription	string	`json:"plain_text_no_images_description"`
	RsvpCloseOffset		int	`json:"rsvp_close_offset"`
	RsvpLimit			int	`json:"rsvp_limit"`
	RsvpOpenOffset		int	`json:"rsvp_open_offset"`
	RsvpRules			RsvpRules	`json:"rsvp_rules"`
	// RsvpSample
	Rsvpable			bool	`json:"rsvpable"`
	RsvpableAfterJoin	bool	`json:"rsvpable_after_join"`
	Saved				bool	`json:"saved"`
	// Self
	// Series
	ShortLink			string	`json:"short_link"`
	SimpleHtmlDescription	string	`json:"simple_html_description"`
	Status			string	`json:"status"`
	// Survey Questions
	Time			int	`json:"time"`
	Updated			int	`json:"updated"`
	UtcOffset		int	`json:"utc_offset"`
	Venue			EventVenue	`json:"venue"`
	VenueVisibility	string	`json:"venue_visibility"`
	Visibility		string	`json:"visibility"`
	WaitlistCount	int		`json:"waitlist_count"`
	Why				string	`json:"why"`
	YesRsvpCount	int		`json:"yes_rsvp_count"`
}


// Event represents a Meetup event
type Member struct {
	Created       int         `json:"created"`
	Joined      int         `json:"joined"`
	ID            int      `json:"id"`
	Name          string      `json:"name"`
	Bio			  string	  `json:"bio"`
	Birthday	  *ExplodedDate	`json:"birthday"`
	City		  string		`json:"city"`
	Country		  string		`json:"country"`
	Email			string		`json:"email"`
	Gender			string		`json:"gender"`
	LastEvent		*Event		`json:"last_event"`
	Lat				float64		`json:"lat"`
	LocalizedCountryName	string	`json:"localized_country_name"`
	Lon				float64		`json:"lon"`
	Memberships		*Memberships	`json:"memberships"`
	Status        string      `json:"status"`
	NextEvent		*Event		`json:"next_event"`
	OtherServices	*[]OtherService	`json:"other_services"`
	Photo			*Photo	`json:"photo"`
	Privacy			*Privacy	`json:"privacy"`
	State			string	`json:"state"`
	Stats			*Stats	`json:"stats"`
	Topics			*[]Topic	`json:"topics"`

}

// GetSelf gets the logged in users data
// Meetup docs: https://www.meetup.com/meetup_api/docs/members/:id
func (c *Client) GetSelf() (*Member, error) {
	fields := "topics,privacy,gender,last_event,memberships,next_event,other_services,self,stats"
	url := fmt.Sprintf("%v/members/self?fields=%v", c.BaseURL,fields)
	log.Println("\n\n")
	log.Println(url)
	log.Println("\n\n")
	req, err := c.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var member *Member
	err = c.Do(req, &member)
	if err != nil {
		return nil, err
	}

	return member, nil
}

// GetMyEvents gets the logged in users data
// Meetup docs: https://www.meetup.com/meetup_api/docs/members/:id
func (c *Client) GetMyEvents() (*[]Rsvps, error) {
	fields := "comment_count,description_images,event_hosts,featured,featured_photo,how_to_find_us,plain_text_description,plain_text_no_images_description,rsvp_rules,rsvpable,rsvpable_after_join,saved,short_link,simple_html_description,venue_visibility"
	url := fmt.Sprintf("%v/self/events?fields=%v", c.BaseURL,fields)
	log.Println("\n\n")
	log.Println(url)
	log.Println("\n\n")
	req, err := c.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var rsvps *[]Rsvps
	err = c.Do(req, &rsvps)
	if err != nil {
		return nil, err
	}

	return rsvps, nil
}
