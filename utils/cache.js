// const cheerio = require('cheerio')
// const fs = require('fs');

// var data = fs.readFileSync('dist/index.html', 'utf-8');
// const $ = cheerio.load(data)

// const now = new Date();

// $("meta[http-equiv=last-modified]").attr("content", now.toUTCString());

// fs.writeFileSync("dist/index.html", $.html());

const fs = require('fs');

fs.writeFileSync('.etag', Math.random());