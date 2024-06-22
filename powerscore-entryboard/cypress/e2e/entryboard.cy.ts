describe('Did it reach out', () => {
    it('Grabs data from a remote server', () => {
      cy.visit('/')
      cy.intercept({ url: '/teams' }).as('request')
      cy.wait('@request')
      cy.get('.assignment').should('contain.text', 'Awesome Team')
    })
  })