import { createStore } from "vuex";
import axios from "axios";

export default createStore({
  state: {
    planSelected: "FaleMais 30",
    plansCost: [],
    dddCodes: [],
    isLoading: false,
    showTable: false,
    originCodes: [],
    destinationCodes: [],
    origin: 0,
    destination: 0,
    minutes: 0,
  },

  getters: {
    getIsLoading(state) {
      return state.isLoading;
    },
    getShowTable(state) {
      return state.showTable;
    },
    getDDDCodes(state) {
      return state.dddCodes;
    },
    getOrigin(state) {
      return state.origin;
    },
    getDestination(state) {
      return state.destination;
    },
    getPlansCost(state) {
      return state.plansCost;
    },
    getPlanSelected(state) {
      return state.planSelected;
    },
    getMinutes(state) {
      return state.minutes;
    },
  },
  mutations: {
    SET_LOADING(state) {
      state.isLoading = !state.isLoading;
    },
    SET_SHOWTABLE(state) {
      state.showTable = !state.showTable;
    },
    SET_DDDCODES(state, payload) {
      state.dddCodes = payload;
    },
    SET_PLAN_COST(state, payload) {
      state.plansCost = payload;
    },
    SET_PlAN_SELECTED(state, payload) {
      state.planSelected = payload;
    },
    SET_ORIGIN(state, payload) {
      state.origin = payload;
    },
    SET_DESTINATION(state, payload) {
      state.destination = payload;
    },
    SET_ORIGIN_CODES(state, payload) {
      state.originCodes = payload;
    },
    SET_DESTINATION_CODES(state, payload) {
      state.destinationCodes = payload;
    },
    SET_MIN(state, payload) {
      state.minutes = payload;
    },
  },
  actions: {
    fetchDDDCodes(context) {
      context.commit("SET_LOADING");
      axios
        .get("/api/v1/dddcost")
        .then((response) => {
          const results = response.data;
          context.commit("SET_LOADING");
          context.commit("SET_DDDCODES", results);
          context.commit("SET_ORIGIN", results[0].origin);
          context.commit("SET_DESTINATION", results[0].destination);
          const origins = results.map((item: any) => item.origin);
          const destinations = results.map((item: any) => item.destination);
          context.commit("SET_ORIGIN_CODES", origins);
          context.commit("SET_DESTINATION_CODES", destinations);
        })
        .catch((error) => {
          context.commit("SET_LOADING");
          console.log(error);
        });
    },
    fetchPlanCost(context, payload) {
      if (!this.state.isLoading) {
        const data = {
          origin: payload.origin,
          destination: payload.destination,
          minutes: payload.minutes,
        };
        axios
          .post("/api/v1/dddcost", data)
          .then((response) => {
            const result = response.data;
            if (this.state.isLoading) {
              context.commit("SET_LOADING");
            }
            if (!this.state.showTable) {
              context.commit("SET_SHOWTABLE");
            }
            context.commit("SET_PLAN_COST", result);
          })
          .catch((error) => {
            context.commit("SET_LOADING");
            console.log(error);
          });
      }
    },
    setOrigin(context, payload) {
      context.commit("SET_ORIGIN", payload);
    },
    setDestination(context, payload) {
      context.commit("SET_DESTINATION", payload);
    },
    setOriginCodes(context, payload) {
      const result = payload.map((item: any) => item.origin);
      context.commit("SET_ORIGIN_CODES", result);
    },
    setDestinationCodes(context, payload) {
      const result = payload.map((item: any) => item.destination);
      context.commit("SET_DESTINATION_CODES", result);
    },
    setMin(context, payload) {
      context.commit("SET_MIN", payload);
    },
    setPlanSelected(context, payload) {
      context.commit("SET_PlAN_SELECTED", payload);
    },
  },
  modules: {},
});
