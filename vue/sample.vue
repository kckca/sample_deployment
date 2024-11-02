<template>
    <v-app>
    <v-container>
      <v-row justify="center">
        <v-col cols="12" md="8" lg="6">
          <v-card>
            <v-card-title class="headline">Submit Topic Information</v-card-title>
            <v-card-text>
              <v-form ref="form">
                <v-text-field v-model="topic" label="Topic Description" required></v-text-field>
                <v-text-field v-model="timing" label="Timing" required></v-text-field>
                <v-text-field v-model="description" label="Description" required></v-text-field>
                <v-btn color="secondary" @click="addRecord">Add</v-btn>
              </v-form>
              <v-data-table :headers="headers" :items="records" class="elevation-1">
                <template v-slot:items="props">
                  <td>{{ props.item.topic }}</td>
                  <td>{{ props.item.timing }}</td>
                  <td>{{ props.item.description }}</td>
                  <td>
                    <v-btn icon @click="removeRecord(props.item)">
                      <v-icon>mdi-delete</v-icon>
                    </v-btn>
                  </td>
                </template>
              </v-data-table>
            </v-card-text>
            <v-card-actions>
              <v-btn color="primary" @click="submitRecords">Submit</v-btn>
            </v-card-actions>
            <v-alert v-if="responseMessage" type="success" dismissible>{{ responseMessage }}</v-alert>
          </v-card>
        </v-col>
      </v-row>
    </v-container>
  </v-app>
</template>
<script>
new Vue({
  el: '#app',
  vuetify: new Vuetify(),
  data: {
    topic: '',
    timing: '',
    description: '',
    records: [],
    headers: [
      { text: 'Topic', value: 'topic' },
      { text: 'Timing', value: 'timing' },
      { text: 'Description', value: 'description' },
      { text: 'Actions', value: 'actions', sortable: false }
    ],
    responseMessage: ''
  },
  methods: {
    addRecord() {
      if (this.topic && this.timing && this.description) {
        this.records.push({
          topic: this.topic,
          timing: this.timing,
          description: this.description
        });
        this.topic = '';
        this.timing = '';
        this.description = '';
      }
    },
    removeRecord(item) {
      const index = this.records.indexOf(item);
      if (index > -1) {
        this.records.splice(index, 1);
      }
    },
    submitRecords() {
      axios.post('/submit', this.records)
        .then(response => {
          this.responseMessage = response.data.message;
          this.records = [];
          setTimeout(() => {
            this.responseMessage = '';
          }, 3000);
        })
        .catch(error => {
          console.error('There was an error!', error);
        });
    }
  }
});
</script>
