describe('Did it reach out', () => {
    it('Attempts to connect to remote server', () => {
      cy.visit('/')
      cy.contains('h1', 'You did it!')
    })
  })