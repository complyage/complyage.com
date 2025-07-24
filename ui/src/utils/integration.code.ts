//||------------------------------------------------------------------------------------------------||
//|| Function :: authLogout
//||------------------------------------------------------------------------------------------------||

      export function integrationCode( publicKey: string ): string {
            const integrationCode = `
<script>
      window._myIntegration = {
            publicKey: "${publicKey}"
      };
</script>
<script src="https://example.com/integration.js"></script>`;

            return integrationCode.trim();
      }
