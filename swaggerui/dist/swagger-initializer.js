window.onload = function() {
  //<editor-fold desc="Changeable Configuration Block">

  const queryParams = new URLSearchParams(window.location.search);
  const apiUrl = queryParams.get('url') || "https://petstore.swagger.io/v2/swagger.json";
  // the following lines will be replaced by docker/configurator, when it runs in a docker-container
  window.ui = SwaggerUIBundle({
    url: apiUrl,
    dom_id: '#swagger-ui',
    deepLinking: true,
    presets: [
      SwaggerUIBundle.presets.apis,
      SwaggerUIStandalonePreset
    ],
    plugins: [
      SwaggerUIBundle.plugins.DownloadUrl
    ],
    layout: "StandaloneLayout"
  });

  //</editor-fold>
};
