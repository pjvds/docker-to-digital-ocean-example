status_code=$(curl --write-out %{http_code} "http://counter.born2code.net/foo")
if [ $status_code -ne 200 ]; then
  fail "Deploy failure: site returned $status_code, 200 was expected";
fi
