#include <curl/curl.h>
int main() {
   char *s1 = "http://speedtest.lax.hivelocity.net";
   CURL *r1 = curl_easy_init();
   curl_easy_setopt(r1, CURLOPT_URL, s1);
   curl_easy_perform(r1);
}
