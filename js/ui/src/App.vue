<template>
  <v-app>
    <v-app-bar
        app
        color="primary"
        dark
    >
      <div class="d-flex align-center">
        <h2>UrFU Abiturient</h2>
      </div>

      <v-spacer></v-spacer>

      <v-btn
          href="https://github.com/maxifom/urfu-abiturient/"
          target="_blank"
          text
      >
        <span class="mr-2">Github</span>
        <v-icon>mdi-open-in-new</v-icon>
      </v-btn>
    </v-app-bar>

    <v-main>
      <v-alert border="left" style="max-width: 1000px" class="mt-4 ml-4" colored-border v-if="error" type="error">
        {{ error }}
      </v-alert>
      <v-container class="mt-4">
        <v-flex>
          <v-row class="ml-1">
            <v-form>
              <v-layout row wrap>
                <v-autocomplete clearable :items="dictionary.program" v-model="program"
                                label="Программа*" :allow-overflow="false" style="max-width: 210px" autocomplete="off"
                >
                </v-autocomplete>
                <v-autocomplete clearable class="ml-8" :items="dictionary.form" v-model="form"
                                label="Форма обучения" multiple :allow-overflow="false" style="max-width: 210px"
                                autocomplete="off"
                >
                  <template v-slot:selection="data">
                    <template v-if="data.index===0">
                      {{ form.length }} формы
                    </template>
                  </template>
                </v-autocomplete>
                <v-autocomplete clearable class="ml-8" :items="dictionary.type" v-model="type_selected"
                                label="Конкурс" multiple :allow-overflow="false" style="max-width: 210px"
                                autocomplete="off"
                >
                  <template v-slot:selection="data">
                    <template v-if="data.index===0">
                      {{ type_selected.length }} конкурса
                    </template>
                  </template>

                </v-autocomplete>
                <v-autocomplete clearable class="ml-8" :items="dictionary.status" v-model="status"
                                label="Статус" multiple :allow-overflow="false" style="max-width: 210px"
                                autocomplete="off"
                >
                  <template v-slot:selection="data">
                    <template v-if="data.index===0">
                      {{ status.length }} статусов
                    </template>
                  </template>
                </v-autocomplete>
                <v-autocomplete clearable class="ml-8" :items="dictionary.basis" v-model="basis"
                                label="Основа" multiple :allow-overflow="false" style="max-width: 210px"
                                autocomplete="off"
                >
                  <template v-slot:selection="data">
                    <template v-if="data.index===0">
                      {{ basis.length }} основ
                    </template>
                  </template>
                </v-autocomplete>
                <v-text-field type="number" style="max-width: 200px;" class="ml-4" v-model="sum.gte" label="Баллы От"
                              autocomplete="off"></v-text-field>
                <v-text-field type="number" style="max-width: 200px;" class="ml-4" v-model="sum.ltee" label="До"
                              autocomplete="off"></v-text-field>
              </v-layout>
            </v-form>
          </v-row>
        </v-flex>

        <v-row class="mt-8">
          <v-checkbox v-model="original_given" label="Оригинал"></v-checkbox>
          <v-checkbox v-model="statement_given" class="ml-4" label="Согласие"></v-checkbox>
        </v-row>
        <v-row class="mt-8">
          <v-text-field type="number" style="max-width: 200px;" v-model="number" label="Номер студента"
                        autocomplete="off"></v-text-field>
          <v-text-field v-model="name" style="max-width: 200px;" class="ml-8" label="ФИО студента"
                        autocomplete="off"></v-text-field>
        </v-row>
        <v-row class="mt-8">
          <v-btn color="primary" @click="query">Поиск</v-btn>
        </v-row>

        <div class="mt-8">
          <v-card>
            <v-card-title>
              <!--              <v-spacer></v-spacer>-->
              <!--              <v-spacer></v-spacer>-->
              <v-autocomplete :items="default_headers" class="mr-16" style="max-width: 200px;" v-model="headers"
                              label="Колонки" multiple :allow-overflow="false"
                              autocomplete="off" item-text="text" return-object
              >
                <template v-slot:selection="data">
                  <template v-if="data.index===0">
                    {{ headers.length }} колонок
                  </template>
                </template>
              </v-autocomplete>
              <v-spacer></v-spacer>
              <v-spacer></v-spacer>
              <v-spacer></v-spacer>
              <v-spacer></v-spacer>
              <v-spacer></v-spacer>
              <v-text-field
                  v-model="search"
                  append-icon="mdi-magnify"
                  label="Поиск"
                  single-line
                  hide-details
              ></v-text-field>
            </v-card-title>
            <v-data-table :search="search" :loading="query_loading" :headers="headers" :items="table"
                          :items-per-page="100" sort-by="sum" sort-desc>
              <template v-slot:item.original_given="{ item }">
                <v-simple-checkbox
                    v-model="item.original_given"
                    disabled
                ></v-simple-checkbox>
              </template>
              <template v-slot:item.statement_given="{ item }">
                <v-simple-checkbox
                    v-model="item.statement_given"
                    disabled
                ></v-simple-checkbox>
              </template>
              <template v-slot:item.index="{ index }">
                {{ index + 1 }}
              </template>
            </v-data-table>
          </v-card>
        </div>
        <v-row class="mt-8 pb-8">
          Last updated &nbsp;
          <timeago :datetime="last_update_time"></timeago>
        </v-row>
      </v-container>
    </v-main>
  </v-app>
</template>

<script>
const axios = require('axios');

export default {
  name: 'App',

  data: () => ({
    dictionary: {
      basis: [],
      form: [],
      program: [],
      type: [],
      status: [],
    },
    last_update_time: null,
    time_loading: false,
    dictionary_loading: false,
    error: null,
    query_loading: false,

    //  inputs:
    name: null,
    program: null,
    form: null,
    type_selected: null,
    status: null,
    number: null,
    sum: {
      lte: null,
      gte: null,
    },
    statement_given: null,
    original_given: null,
    basis: null,

    // end inputs
    search: "",
    default_headers: [
      {
        text: "#",
        value: "index",
        sortable: false,
      },
      {
        text: "Специальность",
        value: "speciality",
        sortable: false,
      },
      {
        text: "Программа",
        value: "program",
        sortable: false,
      },
      {
        text: "Статус",
        value: "status",
        sortable: false,
      },
      {
        text: "Номер",
        value: "number",
        sortable: false,
      },
      {
        text: "Конкурс",
        value: "type",
        sortable: false,
      },

      {
        text: "Согласие",
        value: "statement_given",
        sortable: false,
      },
      {
        text: "Оригинал",
        value: "original_given",
        sortable: false,
      },

      {
        text: "ФИО",
        value: "name",
        sortable: false,
      },
      {
        text: "Форма",
        value: "form",
        sortable: false,
      },
      {
        text: "Основа",
        value: "basis",
        sortable: false,
      },
      {
        text: "Баллы",
        value: "sum",
        align: "end",
        sortable: true,
      }
    ],
    headers: [
      {
        text: "#",
        value: "index",
        sortable: false,
      },
      {
        text: "Статус",
        value: "status",
        sortable: false,
      },
      {
        text: "Номер",
        value: "number",
        sortable: false,
      },
      {
        text: "Конкурс",
        value: "type",
        sortable: false,
      },

      {
        text: "Согласие",
        value: "statement_given",
        sortable: false,
      },
      {
        text: "Оригинал",
        value: "original_given",
        sortable: false,
      },

      {
        text: "ФИО",
        value: "name",
        sortable: false,
      },
      {
        text: "Форма",
        value: "form",
        sortable: false,
      },
      {
        text: "Основа",
        value: "basis",
        sortable: false,
      },
      {
        text: "Баллы",
        value: "sum",
        align: "end",
        sortable: true,
      }
    ],
    table: [],
  }),
  mounted() {
    this.getDictionary();
    this.getLastUpdateTime();
  },
  methods: {
    getDictionary() {
      axios.get("/api/dictionary").then(d => {
        this.error = null;
        this.dictionary = d.data;
      }).catch(e => this.error = e).finally(() => this.dictionary_loading = false);

    },
    getLastUpdateTime() {
      this.time_loading = true;
      axios.get("/api/last_update_time").then(d => {
        this.error = null;
        this.last_update_time = new Date(Date.parse(d.data.last_update_time));
      }).catch(e => this.error = e).finally(() => this.time_loading = false);
    },
    query() {
      this.query_loading = true;
      this.error = null;
      let parsedLte = parseInt(this.sum.lte);
      if (parsedLte !== 0 && !parsedLte) {
        parsedLte = null
      }
      let parsedGte = parseInt(this.sum.gte);
      if (parsedGte !== 0 && !parsedGte) {
        parsedGte = null
      }
      axios.post("/api/query", {
        "basis": this.basis,
        "form": this.form,
        "sum": {
          lte: parsedLte,
          gte: parsedGte,
        },
        "program": this.program,
        "statement_given": this.statement_given,
        "original_given": this.original_given,
        "type": this.type_selected,
        "status": this.status,
        "name": this.name,
        "number": parseInt(this.number) || null,
      }).then(d => {
        this.table = d.data;
      }).catch(e => {
        return this.error = e.response.data.message || e.data.message;
      }).finally(() => this.query_loading = false);
    }
  }
};
</script>

<style>
.v-data-footer__select {
  visibility: hidden;
}
</style>