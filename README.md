# Rails app using React and Go RPC server

## Gems used
* Go server is integrated using the quartz gem
* React uses react-rails and bootstrap-saas

## Add New Relic monitoring
If you want to include New Relic monitoring, first run
```sh
gem install newrelic_rpm
```
if you don't already have the New Relic ruby agent on your system.

Then uncomment the last line of the Gemfile to enable that gem.
Also add your license key to `config/newrelic.yml`

## How to install
Setup rails as follows
```sh
git clone https://github.com/DavidSantia/tagdemo-rails
cd tagdemo-rails
bundle install
```

### Installing Gems on OS X
Some gems require options to build, for example the OpenSSL library path from Homebrew:
```sh
bundle config build.mysql2 "--with-ldflags=-L/usr/local/opt/openssl/lib --with-cppflags=-I/usr/local/opt/openssl/include"
```
Then `bundle install` can work properly

## How to run
In one terminal, setup the test database
```sh
docker build -t demo-mysql ./mysql
```
And run the database container
```sh
docker run --name demo-mysql --rm -p 3306:3306 demo-mysql
```

Make the Go server as follows
```sh
./build-goserver.sh
```

In another terminal, run rails
```sh
cd tagdemo-rails
bin/rails s
```

Go to [localhost:3000](http://localhost:3000)
