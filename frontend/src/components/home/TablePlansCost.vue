<style lang="scss" scopped>
@import "./styles/TablePlansCost";
</style>
<template>
  <div v-show="show" class="div-plans">
    <h1>{{ planName }}</h1>
    <table class="table-plans">
      <thead>
        <tr>
          <th v-for="tHead in tHeads" :key="tHead">{{ tHead }}</th>
        </tr>
      </thead>
      <tbody>
        <tr>
          <td>{{ origin }}</td>
          <td>{{ destination }}</td>
          <td>{{ minutes }}</td>
          <td>{{ planSelected.name }}</td>
          <td>{{ planSelected.with }}</td>
          <td>{{ planSelected.without }}</td>
        </tr>
      </tbody>
    </table>
    <h2>Comparativo com outros planos:</h2>
    <table v-show="show" class="table-plans">
      <thead>
        <tr>
          <th v-for="tHead in tHeads" :key="tHead">{{ tHead }}</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="item in array" :key="item.name">
          <td>{{ origin }}</td>
          <td>{{ destination }}</td>
          <td>{{ minutes }}</td>
          <td>{{ item.name }}</td>
          <td>{{ item.with }}</td>
          <td>{{ item.without }}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>
<script>
import { mapGetters } from "vuex";
//import "./styles/TableComponent.scss";

export default {
  name: "TablePlansCost",
  props: {
    show: Boolean,
  },
  data() {
    return {
      tHeads: [
        "Origem",
        "Destino",
        "Tempo",
        "Plano FaleMais",
        "Com FaleMais",
        "Sem FaleMais",
      ],
      array: [],
      origin: 0,
      destination: 0,
      minutes: 0,
      planName: "",
      planSelected: {},
    };
  },
  computed: {
    ...mapGetters({
      data: "getPlansCost",
      originData: "getOrigin",
      destinationData: "getDestination",
      minutesData: "getMinutes",
      planData: "getPlanSelected",
    }),
  },
  watch: {
    data() {
      this.planSelected = this.data.plan.find(
        (item) => item.name === this.planData
      );
      this.planName = this.planSelected.name;
      const res = this.data.plan.filter((item) => item.name !== this.planData);
      this.array = res.sort((a, b) => b.with - a.with);
      this.origin = this.data.origin;
      this.destination = this.data.destination;
      this.minutes = this.data.minutes;
    },
  },
};
</script>
