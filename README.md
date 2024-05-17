# Slugify

Gem to generate slugs from strings to use in URLs etc..

## Installation

```
gem 'slugify', :git => 'git://github.com/kmatsoukas/slugify-rb.git'
```

Requires `Golang` and `make` installed to compile native extensions.

## Usage

```
require 'slugify'

Slugify.slugify "My Slug"
```

## Development

After checking out the repo, run `bin/setup` to install dependencies. You can also run `bin/console` for an interactive prompt that will allow you to experiment.

To install this gem onto your local machine, run `bundle exec rake install`. To release a new version, update the version number in `version.rb`, and then run `bundle exec rake release`, which will create a git tag for the version, push git commits and the created tag, and push the `.gem` file to [rubygems.org](https://rubygems.org).

## License

The gem is available as open source under the terms of the [MIT License](https://opensource.org/licenses/MIT).
