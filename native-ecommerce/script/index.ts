import { createClient } from "@hey-api/openapi-ts";



try {
  await createClient({
    input: "../go-ecommerce-api/docs/openapi.json",
    output: "./client",
    plugins: [
      "@hey-api/client-axios", 
      {
        name: "@tanstack/react-query",
        infiniteQueryKeys: true, 
        queryKeys: {
          tags: true, 
        },
      },
      "zod",
    ],
  });
  console.log("SDK generated successfully");
} catch (error) {
  console.error("Error generating SDK:", error);
  process.exit(1);
}