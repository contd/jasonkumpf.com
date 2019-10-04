<template>
  <v-layout
    class="d-flex flex-row flex-wrap justify-space-around"
  >
    <v-container>
      <h1 class="display-1">
        About
      </h1>
      <p>
        Born and raised in Cincinnati, OH but now residing just outside of Washington D.C.
      </p>
      <h1 class="display-1">
        Work
      </h1>
      <p>
        An experienced software engineer (20+ years) mostly in web application and infrastructure development. I spent 7 years pursuing my PhD in Electrical Engineering
        building software platforms and solutions for data science research. After teaching as a full time adjunct professor for one year I decided to come back to what
        I am most passionate. I have 2 Bachelors degrees in Information Engineering Technology and Mechanical Engineering Technology from the
        <a href="https://www.uc.edu" target="_blank">University of Cincinnati</a> where I did my graduate work and teaching.  Most of my work has been for software
        development companies but also with non-profits, educational institutions and with small and non-technical companies.  I
        have very strong written and oral communication skills.  I consider myself to be very responsible and always punctual to deadlines.
      </p>
      <p>
        Currently I work as a freelance consultant developing web, mobile and middle-ware apps that bridge the gap between the various open hardware embedded systems
        some of which I have designed and built as an electronics hobbyist. These systems integrate with existing infrastructure as well as new apps.  They often require
        various automated solutions to companies that want full control and customization while leveraging cutting edge technologies.  Many implementations involve
        deploying to multiple cloud providers (VPS/IaaS) providers that require integration with their own private cloud infrastructure and leveraging common design
        patterns and top level security and privacy.  Hybrid cloud solution integrations are usually required for high customization and automation of infrastructure
        on demand. I utilize various tools to develop and deliver programmable infrastructure to streamline the process of deploying reconfigurable development workflows
        and adaptable infrastructure management across multiple service providers.
      </p>
      <p class="d-flex justify-md-space-around">
        <img src="~/assets/aws.svg" alt="AWS" width="100" />
        <img src="~/assets/azure.svg" alt="Azure" width="100" />
        <img src="~/assets/digital-ocean.svg" alt="Digital Ocean" width="100" />
        <img src="~/assets/docker.svg" alt="Docker" width="100" />
        <img src="~/assets/google-cloud-platform.svg" alt="GCP" width="100" />
        <img src="~/assets/kubernetes.svg" alt="Kubernetes" width="100" />
      </p>
      <p>
        Some of the types of systems designed and implemented are deployed to
        <a href="https://aws.amazon.com" target="_blank">AWS</a>,
        <a href="https://azure.microsoft.com" target="_blank">Azure</a>,
        <a href="https://digitalocean.com" target="_blank">Digital Ocean</a> and
        <a href="https://cloud.google.com" target="_blank">Google Cloud Platform</a> in
        <a href="https://kubernetes.io" target="_blank">Kubernetes</a>
        clusters so client apps can be deployed to any of the aforementioned service providers as well as internally. Spreading applications and continuous development
        workflows across multiple cloud service providers and internally makes them highly available and virtually zero potential for outages or downtime.
      </p>
      <p class="d-flex justify-md-space-around">
        <img src="~/assets/nodejs.svg" alt="Node.js" width="100" />
        <img src="~/assets/go-logo.svg" alt="Go" width="100" />
        <img src="~/assets/python.svg" alt="Python" width="100" />
        <img src="~/assets/php.svg" alt="Php" width="100" />
        <img src="~/assets/perl.svg" alt="Perl" width="100" />
        <img src="~/assets/vue.svg" alt="Vue.js" width="100" />
        <img src="~/assets/react.svg" alt="React" width="100" />
        <img src="~/assets/preact.svg" alt="Preact" width="100" />
      </p>
      <p>
        My main skills are development of web services
        (<a href="https://microservices.io" target="_blank">microservices</a>) in
        <a href="https://nodejs.org" target="_blank">Node.js</a>,
        <a href="https://golang.org" target="_blank">Go</a>,
        <a href="https://python.org" target="_blank">Python</a>,
        <a href="https://www.php.net" target="_blank">Php</a> or
        <a href="https://www.perl.org" target="_blank">Perl</a>. Most projects involve developing container
        (<a href="https://docker.com" target="_blank">Docker</a>) based services that work with
        <a href="https://aws.amazon.com" target="_blank">AWS</a> ECS or any
        <a href="https://kubernetes.io" target="_blank">Kubernetes</a>
        cluster. I also have very strong frontend web development skills with various frameworks
        (<a href="https://vuejs.org" target="_blank">Vue.js</a>,
        <a href="https://www.react.org" target="_blank">React</a>,
        <a href="https://preact.org" target="_blank">Preact</a>
        , etc).
      </p>
      <p class="d-flex justify-md-space-around">
        <img src="~/assets/mysql.svg" alt="MySQL" width="100" />
        <img src="~/assets/mssql.svg" alt="Microsoft SQL Server" width="100" />
        <img src="~/assets/postgresql.svg" alt="Postgresql" width="100" />
        <img src="~/assets/sqlite.svg" alt="Sqlite" width="100" />
        <img src="~/assets/mongodb.svg" alt="Mongodb" width="100" />
        <img src="~/assets/redis.svg" alt="Redis" width="100" />
        <img src="~/assets/influxdb.svg" alt="InfluxDB" width="100" />
      </p>
      <p>
        Most applications worked on are a combination of database driven REST APIs. Significant time has been spent working with both relational
        (<a href="https://mysql.org" target="_blank">MySQL</a>,
        <a href="https://www.microsoft.com/en-us/sql-server/" target="_blank">MSSQL</a>,
        <a href="https://www.postgresql.org" target="_blank">Postgresql</a>,
        <a href="https://www.sqlite.org" target="_blank">Sqlite</a>)
        and nosql
        (<a href="https://www.mongodb.com" target="_blank">Mongodb</a>,
        <a href="https://redis.io/" target="_blank">Redis</a>,
        <a href="https://www.influxdata.com/" target="_blank">InfluxDB</a>)
        databases and am very fluent writing and optimizing SQL.
      </p>
    </v-container>
  </v-layout>
</template>

<script>
export default {
  head () {
    return {
      title: 'About'
    }
  }
}
</script>
