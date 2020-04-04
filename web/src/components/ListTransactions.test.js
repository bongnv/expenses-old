import ListTransactions from "./ListTransactions";
import { shallowMount } from "@vue/test-utils";

describe("@components/list-transactions", () => {
  it("exports a valid component", () => {
    expect(shallowMount(ListTransactions).exists()).toBe(true);
  });
});
