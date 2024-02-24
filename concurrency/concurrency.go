package concurrency

type WebsiteChecker func(string) bool

type result struct {
	urlName  string
	urlAvail bool
}

// Accepts: wc WebsiteChecker
// Accpets: urls slice of strings
// Returns: map of string to boolean
// determine whether website is accessible
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			resultChannel <- result{u, wc(u)} // result is sent to resultsChannel?
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		r := <-resultChannel // create r which recieves on resultsChannel?
		results[r.urlName] = r.urlAvail
	}

	return results
}
