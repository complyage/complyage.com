//||------------------------------------------------------------------------------------------------||
//|| Function :: authLogout
//||------------------------------------------------------------------------------------------------||

      export function integrationCode( clientId: string ): string {
            const integrationCode = `<script src="https://gate.complyage.com//v1/complyage.js?client_id=${clientId}"></script>`;
            return integrationCode.trim();
      }
