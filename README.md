# flettsgetmarried

## Web

```bash
cd website

# install dependencies
npm install

# serve with hot reload at localhost:80
npm run dev

# build as SPA
nuxt build --spa

# deploy
cd dist
aws s3 sync . s3://flettsgetmarried.com
```

## API

```bash
./build.sh
serverless deploy

curl --location --request POST 'https://api.flettsgetmarried.com/rsvp' \
--header 'Content-Type: application/json' \
--data-raw '{
    "contact": "0429235371",
    "diet": "None",
    "music": "Opeth, obviously",
    "attendees": ["Ryan Flett"],
    "absentees": ["Kara Hazelman"]
}'
```

# TODO

- Allow people to upload photos
- Make all the photos visible on the site
