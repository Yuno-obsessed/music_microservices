package consts

type BucketName string

const (
	// Bucket with customer's images
	CustomerBucket BucketName = "customer.bucket"
	// Bucket with artist's images etc
	ArtistBucket BucketName = "artist.bucket"
	// Bucket with band's images etc
	BandBucket BucketName = "band.bucket"
	// Bucket with event videos, images
	EventBucket BucketName = "event.bucket"
	// Bucket with QR codes or pdfs of tickets
	TicketBucket BucketName = "ticket.bucket"
)

func GetBucketNames() []BucketName {
	return []BucketName{
		CustomerBucket,
		ArtistBucket,
		BandBucket,
		EventBucket,
		TicketBucket,
	}
}
