import { createStore } from "vuex";
import axios from "axios";

export default createStore({
  state: {
    planCost: {},
    dddCodes: [],
      isLoading: false,
    origin: 0,
    destination: 0,
  },
  getters: {
    getDDDCodes(state) {
      return state.dddCodes;
    },
  },
  mutations: {
      SET_LOADING(state) {
          state.isLoading = !state.isLoading;
      },
      SET_DDDCODES(state, payload) {
          state.dddCodes = payload;
      },
      SET_PLAN_COST(state, payload) {
          state.planCost = payload;
      },
      SET_ORIGIN(state, payload) {
          state.origin = payload;
      },
      SET_DESTINATION(state, payload) {
          state.destination = payload;
      }
  },
  actions: {
    fetchDDDCodes(context) {
      context.commit("SET_LOADING");
      axios
          .get("https://localhost:8080/api/v1/dddcost")
          .then((response) => {
            const { results } = response.data;
            context.commit("SET_LOADING");
            context.commit("SET_DDDCODES", results);
          })
          .catch((error) => {
            context.commit("SET_LOADING");
            console.log(error);
          });
    },
      fetchPlanCost(context, payload) {
          context.commit("SET_LOADING");
          const requestOptions = {
              method: "POST",
              headers: { "Content-Type": "application/json" },
              body: JSON.stringify({ origin: payload.origin, destination: payload.destination }),
          };

          axios
              .post("https://localhost:8080/api/v1/dddcost", requestOptions)
              .then((response) => {
                  const { results } = response.data;
                  context.commit("SET_LOADING");
                  context.commit("SET_PLAN_COST", results);
              })
              .catch((error) => {
                  context.commit("SET_LOADING");
                  console.log(error);
              });
      },
      setOrigin(context, payload) {
            context.commit("SET_ORIGIN", payload);
      },
        setDestination(context, payload) {
            context.commit("SET_DESTINATION", payload);
    }
  },
  modules: {},
});
