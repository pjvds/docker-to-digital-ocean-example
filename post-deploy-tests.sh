status_code=$(curl --write-out %{http_code} --silent "http://counter.born2code.net/foo")
if [ $status_code -ne 200 ]; then
  fail "Deploy failure: site returned $status_code, 200 was expected";
fi
success 'Post deploy tests succeeded'
