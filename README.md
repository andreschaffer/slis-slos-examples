# SLIs and SLOs examples

# Motivation
How reliable is the service you provide to your users? What would be a _good enough_ reliability level for your service? And what would change in how you work with software development if reliability was a feature?  
There is a lot of material out there about SLIs and SLOs, but to me these three simple questions reveal the essence of it.  

In this project I will share a practical example of working with these concepts. I will try and keep it short and to the point, and recommend reading [Google's SRE book](https://sre.google/sre-book/service-level-objectives) if one wants to learn the subject more in depth.

# Terminology
SLIs stand for Service Level Indicators. They are simply the indicators or _measurements_ of your service that speak for its reliability. They will help you with the first question “How reliable is the service you provide to your users?”.

SLOs stand for Service Level Objectives. They are simply your SLIs + objectives or _targets_ for your service’s reliability. They clearly relate to the second question “What would be a _good enough_ reliability level for your service?”, but also to the third one “What would change in how you work with software development if reliability was a feature?”. Once a reasonable reliability target is defined, running below it should spur action to get back on track, as no one wants to provide an unreliable experience to their users. Or in other words, reliability now becomes an unambiguous feature, and should be fixed if broken.

SLAs stand for Service Level Agreements. They are simply your SLOs + agreements or _consequences_ of breaking the SLOs. These are signed between providers and customers, and the consequences are usually monetary. They directly speak to the third question “What would change in how you work with software development if reliability was a feature?”, as no one wants to lose money. I just wanted to briefly touch on the concept of SLAs since it’s common in the reliability and business world, but won’t go further on it in this project.

# Working with it
Now with the terms covered, it’s time to look at defining SLIs and SLOs for our services.

## SLIs
The types of measurements that can speak for the reliability of your service may vary depending on its nature. In most cases though, they will fall into one of the categories from Google's SLI Menu.

TODO

## SLOs
TODO

## Alerting
TODO

## Where to start in your organization?
My first advice, as with anything else, is to remain intentional. Do not care too much about measuring, but measure what you care about. It's easy to fall for measuring a lot of things just because one can. Understand what really matters to your users, and focus on those.

My second advice is to start pragmatically and from your external users' perspective. Look at the "edge" of your service, where your external users interact directly with, and define your SLIs and SLOs there. As your service's system grows in complexity and dependencies, add SLIs and SLOs to the dependencies too. It will be tricky to maintain SLIs and SLOs at the "edge" and to understand where things went wrong if the dependencies do not take their part in it too.

# Running example
TODO

# Contributing
If you would like to help making this project better, see the [CONTRIBUTING.md](CONTRIBUTING.md).  

# Maintainers

Send any other comments, flowers and suggestions to [André Schaffer](https://github.com/andreschaffer).

# License
This project is distributed under the [MIT License](LICENSE).
