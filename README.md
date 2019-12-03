# Advent of Code - 2019 - Golang Edition

# Auto Answering

POST to `https://adventofcode/{year}/day/{day}/answer`

with form data:

- `level`: `1` for silver, `2` for gold
- `answer`: answer

and `session` cookie

Response looks like (on success):

```html
<!DOCTYPE html>
<html lang="en-us">
  <head>
    <meta charset="utf-8" />
    <title>Day 2 - Advent of Code 2019</title>
    <!--[if lt IE 9]><script src="/static/html5.js"></script><![endif]-->
    <link
      href="//fonts.googleapis.com/css?family=Source+Code+Pro:300&subset=latin,latin-ext"
      rel="stylesheet"
      type="text/css"
    />
    <link rel="stylesheet" type="text/css" href="/static/style.css?22" />
    <link
      rel="stylesheet alternate"
      type="text/css"
      href="/static/highcontrast.css?0"
      title="High Contrast"
    />
    <link rel="shortcut icon" href="/favicon.png" /></head
  ><!--




Oh, hello!  Funny seeing you here.

I appreciate your enthusiasm, but you aren't going to find much down here.
There certainly aren't clues to any of the puzzles.  The best surprises don't
even appear in the source until you unlock them for real.

Please be careful with automated requests; I'm not a massive company, and I can
only take so much traffic.  Please be considerate so that everyone gets to play.

If you're curious about how Advent of Code works, it's running on some custom
Perl code. Other than a few integrations (auth, analytics, ads, social media),
I built the whole thing myself, including the design, animations, prose, and
all of the puzzles.

The puzzles are most of the work; preparing a new calendar and a new set of
puzzles each year takes all of my free time for 4-5 months. A lot of effort
went into building this thing - I hope you're enjoying playing it as much as I
enjoyed making it for you!

If you'd like to hang out, I'm @ericwastl on Twitter.

- Eric Wastl


















































-->
  <body>
    <header>
      <div>
        <h1 class="title-global"><a href="/">Advent of Code</a></h1>
        <nav>
          <ul>
            <li><a href="/2019/about">[About]</a></li>
            <li><a href="/2019/events">[Events]</a></li>
            <li>
              <a href="https://teespring.com/adventofcode-2019" target="_blank"
                >[Shop]</a
              >
            </li>
            <li><a href="/2019/settings">[Settings]</a></li>
            <li><a href="/2019/auth/logout">[Log Out]</a></li>
          </ul>
        </nav>
        <div class="user">Matt Steele <span class="star-count">3*</span></div>
      </div>
      <div>
        <h1 class="title-event">
          &nbsp;<span class="title-event-wrap">{'year':</span
          ><a href="/2019">2019</a><span class="title-event-wrap">}</span>
        </h1>
        <nav>
          <ul>
            <li><a href="/2019">[Calendar]</a></li>
            <li><a href="/2019/support">[AoC++]</a></li>
            <li><a href="/2019/sponsors">[Sponsors]</a></li>
            <li><a href="/2019/leaderboard">[Leaderboard]</a></li>
            <li><a href="/2019/stats">[Stats]</a></li>
          </ul>
        </nav>
      </div>
    </header>

    <div id="sidebar">
      <div id="sponsor">
        <div class="quiet">
          Our <a href="/2019/sponsors">sponsors</a> help make Advent of Code
          possible:
        </div>
        <div class="sponsor">
          <a
            href="https://smartystreets.com/aoc"
            target="_blank"
            onclick="if(ga)ga('send','event','sponsor','sidebar',this.href);"
            rel="noopener"
            >SmartyStreets</a
          >
          - Ridiculously -- Fast -------- Address ------ Verification --
        </div>
      </div>
    </div>
    <!--/sidebar-->

    <main>
      <article>
        <p>
          That's the right answer! You are
          <span class="day-success">one gold star</span> closer to rescuing
          Santa. <a href="/2019/day/2#part2">[Continue to Part Two]</a>
        </p>
      </article>
    </main>

    <!-- ga -->
    <script>
      (function(i, s, o, g, r, a, m) {
        i['GoogleAnalyticsObject'] = r;
        (i[r] =
          i[r] ||
          function() {
            (i[r].q = i[r].q || []).push(arguments);
          }),
          (i[r].l = 1 * new Date());
        (a = s.createElement(o)), (m = s.getElementsByTagName(o)[0]);
        a.async = 1;
        a.src = g;
        m.parentNode.insertBefore(a, m);
      })(
        window,
        document,
        'script',
        '//www.google-analytics.com/analytics.js',
        'ga'
      );
      ga('create', 'UA-69522494-1', 'auto');
      ga('set', 'anonymizeIp', true);
      ga('send', 'pageview');
    </script>
    <!-- /ga -->
  </body>
</html>
```
