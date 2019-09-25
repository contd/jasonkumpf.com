<template>
  <v-layout
    class="d-flex flex-row flex-wrap justify-space-around"
  >
    <v-container>
      <span class="display-2 mr-2">
        Jason Kumpf
      </span>
      <v-divider />
      {{ basics.email }} |
      <a :href="basics.website" target="_blank">{{ basics.website }}</a>
    </v-container>
    <v-container>
      <h1>
        Work
      </h1>
      <v-divider />
      <v-timeline>
        <v-timeline-item
          v-for="(item, n) in work"
          :key="n"
        >
          <span slot="opposite">
            <a :href="item.website" target="_blank" class="headline">
              {{ item.company }}
            </a>
            <p>
              {{ item.location }}
            </p>
          </span>
          <v-card class="elevation-2">
            <v-card-title class="headline">
              {{ item.position }}
            </v-card-title>
            <v-card-text>
              <p>
                {{ item.startDate }} - {{ item.endDate }}
              </p>
              <p>
                {{ item.summary }}
              </p>
            </v-card-text>
          </v-card>
        </v-timeline-item>
      </v-timeline>
    </v-container>
    <v-container>
      <h1>
        Education
      </h1>
      <v-divider />
      <v-card
        v-for="(edu, i) in education"
        :key="i"
        class="mb-1"
      >
        <v-card-title>
          {{ edu.institution }}
          <span class="ml-4 grey--text subtitle-1">
            {{ edu.startDate }} - {{ edu.endDate }}
          </span>
        </v-card-title>
        <v-card-text>
          {{ edu.area }} ({{ edu.studyType }})
          <ul>
            <li v-for="(extra, k) in edu.extras" :key="k">
              <span class="grey--text subtitle-1">
                {{ extra.name }}:
              </span>
              <span class="ml-2 subtitle-1">
                {{ extra.value }}
              </span>
            </li>
          </ul>
        </v-card-text>
      </v-card>
    </v-container>
    <v-container>
      <v-card
        v-for="(skill, i) in skills"
        :key="i"
        class="mb-1"
      >
        <v-card-title>
          {{ skill.name }} <span class="grey--text subtitle-1">{{ skill.level }}</span>
        </v-card-title>
        <v-card-text>
          <v-chip
            v-for="(kw, j) in skill.keywords"
            :key="j"
            class="ma-2"
            :color="colors[j]"
            text-color="white"
          >
            {{ kw }}
          </v-chip>
        </v-card-text>
      </v-card>
      <v-divider />
    </v-container>
    <v-container>
      <v-card
        v-for="(pub, i) in publications"
        :key="i"
        class="mb-1"
      >
        <v-card-text>
          {{ pub.title }}, {{ pub.authors }}, {{ pub.publication }}, {{ pub.year }}
          <a v-if="pub.website" :href="pub.website">{{ pub.website }}</a>
          <a v-if="pub.website2" :href="pub.website2">{{ pub.website2 }}</a>
          <span v-if="pub.summary">{{ pub.summary }}</span>
        </v-card-text>
      </v-card>
      <v-divider />
    </v-container>
  </v-layout>
</template>

<script>
const resumeJson = require('@/services/resume.json')

export default {
  head () {
    return {
      title: 'Resume'
    }
  },
  filters: {
    jsonString (obj) {
      return JSON.stringify(obj, null, 2)
    }
  },
  data () {
    return {
      sections: ['basics', 'work', 'education', 'skills', 'publications'],
      colors: ['primary', 'secondary', 'red', 'green', 'orange'],
      resume: {},
      basics: {},
      work: [],
      education: [],
      skills: [],
      publications: []
    }
  },
  mounted () {
    this.resume = resumeJson
    this.basics = this.resume.basics
    this.work = this.resume.work
    this.education = this.resume.education
    this.skills = this.resume.skills
    this.publications = this.resume.publications
  }
}
</script>
