# zcrawl

`zcrawl` is an open source software platform to deploy and orchestrate web crawlers and crawling tasks in general. It's written in [Go](https://golang.org/) and one of the goals is to make it as flexible as possible to allow integrations with different languages and third-party services.

In order to avoid any language lock-ins, `zcrawl` will provide enough tools to enhance the process of creating and deploying a web crawler using your favorite language, so it's not Go specific.

We're still in the planning phase and the roadmap is subject to changes. A prototype is in progress and it's being developed as we want to test some of our ideas in a minimal way.

## How to use it?

No instructions are provided at this time, if you're interested feel free to pull the code, build it and see what happens :)

## Is it for me?

The project is targeted to users who want an easy way of deploying web crawlers, without messing up with `crontab` (in case you need to schedule recurrent crawling jobs), plain CSV files (in case you do this straight from the command line), multi-worker environments (when you need to orchestrate a distributed crawling task) and more complex pipelines that might be a combination of all these tasks.

Think about this as a Heroku-like solution where you can deploy text crawlers and orchestrate them to re-train your machine learning models with fresh data, everything in your own infrastructure. This is the type of scenarios we're interested in.

## Roadmap

TBA

## Contact

hello AT zcrawl DOT org