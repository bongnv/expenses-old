import Expenses from './expenses'

describe('@views/expenses', () => {
  it('is a valid view', () => {
    expect(Expenses).toBeAViewComponent()
  })
})
