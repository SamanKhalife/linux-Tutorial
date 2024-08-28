# Some basic info about Cloud Computing model and SRE's 

## Cloud Delivery Models

#### Infrastructure as a Service (IaaS)

IaaS is the on-demand delivery of computing infrastructure, including operating systems, networking, storage, and other infrastructural components. Acting much like a virtual equivalent to physical servers, IaaS relieves cloud users of the need to buy and maintain physical servers while also providing the flexibility to scale and pay for resources as needed. IaaS is a popular option for businesses that wish to leverage the advantages of the cloud and have system administrators who can oversee the installation, configuration, and management of operating systems, development tools, and other underlying infrastructure that they wish to use. However, IaaS is also used by developers, researchers, and others who wish to customize the underlying infrastructure of their computing environment. Given its flexibility, IaaS can support everything from a company’s computing infrastructure to web hosting to big data analysis.

#### Platform as a Service (PaaS)

PaaS provides a computing platform where the underlying infrastructure (such as the operating system and other software) is installed, configured, and maintained by the provider, allowing users to focus their efforts on developing and deploying apps in a tested and standardized environment. PaaS is commonly used by software developers and developer teams as it cuts down on the complexity of setting up and maintaining computer infrastructure, while also supporting collaboration among distributed teams. PaaS can be a good choice for developers who don’t have the need to customize their underlying infrastructure, or those who want to focus their attention on development rather than DevOps and system administration.

#### Software as a Service (SaaS)

SaaS providers are cloud-based applications that users access on demand from the internet without needing to install or maintain the software. Examples include GitHub, Google Docs, Slack, and Adobe Creative Cloud. SaaS applications are popular among businesses and general users given that they’re often easy to adopt, accessible from any device, and have free, premium, and enterprise versions of their applications. Like PaaS, SaaS abstracts away the underlying infrastructure of the software application so that users are only exposed to the interface they interact with.

![img](https://github.com/SamanKhalife/linux-Tutorial/blob/main/IMAGES/7777777777777777777.png)

## Cloud Environments

#### Public Cloud

The public cloud refers to cloud services (such as virtual machines, storage, or applications) offered publicly by a commercial provider to businesses and individuals. Public cloud resources are hosted on the commercial provider’s hardware, which users access through the internet. They are not always suitable for organizations in highly-regulated industries, such as healthcare or finance, as public cloud environments may not comply with industry regulations regarding customer data.

#### Private Cloud

The private cloud refers to cloud services that are owned and managed by the organization that uses them and available only to the organization’s employees and customers. Private clouds allow organizations to exert greater control over their computing environment and their stored data, which can be necessary for organizations in highly-regulated industries. Private clouds are sometimes seen as more secure than public clouds as they are accessed through private networks and enable the organization to directly oversee their cloud security. Public cloud providers sometimes provide their services as applications that can be installed on private clouds, allowing organizations to keep their infrastructure and data on premise while taking advantage of the public cloud’s latest innovations.

#### Hybrid Cloud and Multicloud

Many organizations use a hybrid cloud environment which combines public and private cloud resources to support the organization’s computing needs while maintaining compliance with industry regulation. Multicloud environments are also common, which entail the use of more than one public cloud provider (for example, combining Amazon Web Services and DigitalOcean).

#### Managed cloud

Managed cloud services refer to cloud computing solutions in which a third-party provider takes on the responsibility of managing various aspects of your cloud infrastructure, applications, and services.

#### Micro cloud

Micro clouds are a new class of infrastructure for on-demand computing at the edge. They differ from the internet-of-things (IoT), which uses thousands of single machines or sensors to gather data, yet they perform computing tasks. Micro clouds reuse proven cloud primitives but with the unattended, autonomous and clustering features that resolve typical edge computing challenges.

#### Multi-cloud

Multi-cloud (also referred to as multi cloud or multicloud) is a concept that refers to using multiple clouds from more than one cloud service provider at the same time. The term is also used to refer to the simultaneous running of bare metal, virtualised and containerised workloads.

## SLA(Service Level agreement) / SLI(Service Level Indicators) / SLO(Service Level Objectives)

![img](https://github.com/SamanKhalife/linux-Tutorial/blob/main/IMAGES/6666666666666666666.png)

If the running time and availability of a system cannot be measured, it is very difficult to maintain and operate the system that is already online, which often causes the maintenance and operation team to continue to be in the state of a fire brigade, and when the root cause of the problem is finally found, it may be the development There is a problem with the code written by the team.

Development teams often don't see "stability" as a potential problem if they can't figure out how to measure runtime and availability. This problem has plagued Google for many years, which is why the SRE principle was developed. One of the motivations.

SRE ensures that everyone knows how to measure reliability and what to do when a service fails. This will be detailed to the point that when a problem occurs, from VP or CxO to every relevant employee within the organization, they all have to do what they should do. What each "person" should do is clearly regulated. SRE will communicate with all relevant personnel to determine Service Level Indicators (SLIs) and Service Level Objectives (SLOs).

SLIs define metrics related to the "response time" of a system, such as response time, throughput per second, requests, etc., and often convert this metric into a ratio or average.

SLOs are a time interval obtained after discussions with relevant personnel. It is expected that SLIs can maintain a certain level of figures, such as "what is the level of SLIs every month", which is more internal indicators.

The video also discusses Service Level Agreements (SLAs), even though it's not a number that SREs care about on a daily basis. As an online service provider, SLA is a commitment to customers to ensure that the percentage of continuous operation of the service is usually "negotiated" with customers, and the downtime per year (or month) should not be less than a few minutes.

The concepts of SLI, SLO, and SLA are very similar to the "everything can be measured" mentioned by DevOps, which is one of the reasons why it is said that class SRE implements DevOps.

#### Uptime

| Availability % | Downtime per year | Downtime per month | Downtime per Week |
| -------------- | ----------------- | ------------------ | ----------------- |
| 90%            | 36.5 days         | 72 hours           | 16.8 hours        |
| 95%            | 18.25 days        | 36 hours           | 8.4 hours         |
| 98%            | 7.30 days         | 14.4 hours         | 3.36 hours        |
| 99%            | 3.65 days         | 7.20 hours         | 1.68 hours        |
| 99.5%          | 1.83 days         | 3.60 hours         | 50.4 minutes      |
| 99.8%          | 17.52 hours       | 86.23 minutes      | 20.16 minutes     |
| 99.9%          | 8.76 hours        | 43.2 minutes       | 10.1 minutes      |
| 99.95%         | 4.38 hours        | 21.56 minutes      | 5.04 minutes      |
| 99.99%         | 52.6 minutes      | 4.32 minutes       | 1.01 minutes      |
| 99.999%        | 5.26 minutes      | 25.9 seconds       | 6.05 seconds      |
| 99.9999%       | 31.5 seconds      | 2.59 seconds       | 0.605 seconds     |

## Mean Time To Recover (MTTR)

MTTR is the average time that a device will take to recover from any failure.

## Mean Time Between Failures (MTBF)

MTBF is the predicted elapsed time between inherent failures of a mechanical or electronic system, during normal system operation. MTBF can be calculated as the arithmetic mean (average) time between failures of a system. The term is used for repairable systems, while mean time to failure (MTTF) denotes the expected time to failure for a non-repairable system.

## Mean Time To Failure (MTTF)

MTTF denotes the expected time to failure for a non-repairable system.


## Critical incident response time (CIRT)

Critical incident response time (CIRT) is a new, significantly more accurate method to evaluate operations performance. CIRT focuses on the incidents that are most likely to impact business by culling noise from incoming signals using the following techniques:

Real business-impacting (or potentially impacting) incidents are very rarely low urgency, so rule out all low-urgency incidents.

Real business-impacting incidents are very rarely (if ever) auto-resolved by monitoring tools without the need for human intervention, so rule out incidents that were not resolved by a human.

Short, bursting, and transient incidents that are resolved within 120 seconds are highly unlikely to be real business-impacting incidents, so rule them out.

Incidents that go unnoticed, tabled, or ignored (not acknowledged, not resolved) for a very long time are rarely business-impacting; rule them out. Note: 

This threshold can be a statistically derived number that is customer-specific (e.g., two standard deviations above the mean) to avoid using an arbitrary number.

Individual, ungrouped incidents generated by separate alerts are not representative of the larger business-impacting incident. Therefore, simulate incident 

groupings with a very conservative threshold, e.g., two minutes, to calculate response time.

## books that you need to read as sre

- [incident-management](https://www.atlassian.com/incident-management/get-the-handbook)
- [Building Secure & Reliable Systems](https://sre.google/books/)
- [The Site Reliability Workbook](https://sre.google/books/)
- [Site Reliability Engineering](https://sre.google/books/)

