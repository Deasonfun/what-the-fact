<template>
  <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0-alpha3/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-KK94CHFLLe+nY2dmCWGMq91rCGa5gtU4mk92HdvYe+M/SXH301p5ILy+dN9+nJOZ" crossorigin="anonymous">
  <div class="fact-card-container">
    <div class="card">
      <div class="card-header">
        Fact:
      </div>
      <div class="card-body">
        <p class="card-text">{{ this.fact }}</p>
        <button @click="NextFact" class="btn btn-success">Next</button>
      </div>
    </div>
  </div>
</template>

<script>
  import axios from 'axios';

  export default {
    name: 'what-the-fact',
    components: {
    },
    data() {
      return {
        fact: ''
      }
    },
    methods: {
      NextFact() {
        axios.get('http://localhost:8080/test')
          .then(res => this.fact = res.data.fact)
      }
    },
    async mounted() {
      axios.get('http://localhost:8080/test')
        .then(res => this.fact = res.data.fact)
    }
   
  }
</script>

<style>
  #app {
    font-family: Avenir, Helvetica, Arial, sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    text-align: center;
    color: #2c3e50;
    margin-top: 60px;
  }
  .fact-card-container {
    position: fixed;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    height: 40rem;
  }
  .card {
    width: 20rem;
    height: 100%;
  }
</style>
