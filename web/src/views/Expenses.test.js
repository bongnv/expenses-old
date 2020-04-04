import Expenses from "./Expenses";
import { shallowMount, createLocalVue } from "@vue/test-utils";
import Vuetify from "vuetify";

describe("@components/Expenses", () => {
  it("exports a valid component", () => {
    const vuetify = new Vuetify();
    vuetify.breakpoint = {
      lgAndUp: true
    };

    const localVue = createLocalVue();
    const wrapper = shallowMount(Expenses, {
      localVue,
      vuetify
    });

    expect(wrapper.exists()).toBe(true);
  });
});
