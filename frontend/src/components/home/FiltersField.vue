<style lang="scss" scopped>
@import "./styles/FiltersField";
</style>
<template>
  <div>
    <div class="div-filters">
      <div class="filter">
        <FilterSelect
          :array="originCodes"
          :label="labelFilterOrigin"
          @update="origin = $event"
        />
        <FilterSelect
          :array="destinationCodes"
          :label="labelFilterDestination"
          @update="destination = $event"
        />
        <FilterSelect
          :array="plansOptions"
          :label="labelFilterPlan"
          @update="plan = $event"
        />
        <InputNumber :label="labelMinutes" @updateValue="minutes = $event" />
        <BtnCalc :labelBtn="labelBtnCalc" @clickBtn="calculate()" />
      </div>
    </div>
  </div>
</template>
<script>
import FilterSelect from "../shared/FilterSelect";
import BtnCalc from "../shared/Btn";
import { mapActions, mapGetters, mapState } from "vuex";
import InputNumber from "@/components/home/InputNumber";

export default {
  name: "FiltersField",
  components: {
    InputNumber,
    FilterSelect,
    BtnCalc,
  },
  data() {
    return {
      labelBtnCalc: "Calcular",
      labelFilterOrigin: "Origem: ",
      labelFilterDestination: "Destino: ",
      labelFilterPlan: "Plano: ",
      labelMinutes: "Minutos: ",
      plansOptions: ["FaleMais 30", "FaleMais 60", "FaleMais 120"],
      minutes: 0,
      origin: 0,
      destination: 0,
      plan: "",
    };
  },
  computed: {
    ...mapState({
      originCodes: (state) => state.originCodes,
      destinationCodes: (state) => state.destinationCodes,
      originData: (state) => state.origin,
      destinationData: (state) => state.destination,
    }),
  },
  methods: {
    ...mapActions([
      "fetchDDDCodes",
      "fetchPlanCost",
      "setOrigin",
      "setDestination",
      "setOriginCodes",
      "setDestinationCodes",
      "setMin",
      "setPlanSelected",
    ]),
    ...mapGetters(["getDDDCodes", "getOrigin", "getDestination"]),
    calculate() {
      let origin = this.origin;
      let destination = this.destination;
      const minutes = this.minutes;
      console.log(origin, destination, minutes);
      if (minutes <= 0) {
        alert("O valor de minutos não pode ser negativo ou igual a zero");
        return;
      }
      if (origin === destination) {
        alert("O DDD de origem e destino não podem ser iguais");
        return;
      }
      this.setPlanSelected(this.plan);
      this.fetchPlanCost({ origin, destination, minutes });
    },
  },
  watch: {
    originCodes() {
      this.origin = this.originCodes[0];
      this.plan = "FaleMais 30";
    },
    destinationCodes() {
      this.destination = this.destinationCodes[0];
    },
  },
  created() {
    this.fetchDDDCodes();
  },
};
</script>
