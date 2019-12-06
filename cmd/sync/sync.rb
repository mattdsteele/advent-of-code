require 'net/http'
require 'cgi'
require 'json'


def status
    puts "Checking submitted answers..."
    year = 2019
    user_id = 31777
    url = "https://adventofcode.com/#{year}/leaderboard/private/view/#{user_id}.json"

    uri = URI(url)
    http = Net::HTTP.new(uri.host, uri.port)
    http.use_ssl = true
    req = Net::HTTP::Get.new(uri.request_uri)
    cookie = CGI::Cookie.new('session', ENV['AOC_SESSION_COOKIE'])
    req['Cookie'] = cookie.to_s
    res = http.request(req)

    parsed = JSON.parse(res.body)
    parsed["members"]["#{user_id}"]["completion_day_level"]
end

def completed_days input
    input.keys.select {|day| input[day].keys.count == 2}
end

def build
    puts "Compiling all days..."
    `rm -rf bin`
    `mkdir bin`
    `go build -o bin ./src/*`
end

def answered_days
    vals = `ls bin`
    vals.split ' '
end

def silver day, results
    results[day] != nil
end

def gold day, results
    return false if results[day] == nil
    results[day]["2"] != nil
end

def has_gold_answer output
    output.length == 2
end
def submit_day day, results
    silver_correct = silver day, results
    gold_correct = gold day, results
    puts "Day #{day} silver: #{silver_correct}, gold: #{gold_correct}"

    out = `./bin/#{day}`
    res = out.split("\n")

    if !silver_correct
        silver_answer = res[0]
        submit day, 1, silver_answer
    end

    if !gold_correct && has_gold_answer(res)
        gold_answer = res[1]
        submit day, 2, gold_answer
    end

    if res.length == 2 then
        puts "Gold: #{res[1]}"
    end
end

def submit day, level, answer
    puts "submitting #{day} level #{level} answer #{answer}"
    res = `go run ./cmd/submit #{day} #{level} #{answer}`
    puts res
end

results = status
completed = completed_days(results)
build
answered = answered_days
answered_not_submitted = answered.select{|day| !completed.include? day}

answered_not_submitted.each do |day|
    puts "Submitting new answer(s) for day #{day}"
    submit_day day, results
end
puts "Done!"