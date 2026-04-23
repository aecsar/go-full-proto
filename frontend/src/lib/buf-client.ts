import { GreetService } from "@pb/greet_pb";

import { createClient } from "@connectrpc/connect";
import { createConnectTransport } from "@connectrpc/connect-web";
import { createValidateInterceptor } from "@connectrpc/validate";

const transport = createConnectTransport({
  baseUrl: "http://localhost:3000",
  interceptors: [createValidateInterceptor()],
});

export const greetClient = createClient(GreetService, transport);
