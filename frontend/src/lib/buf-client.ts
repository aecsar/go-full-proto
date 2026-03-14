import { GreetService } from "@pb/greet_pb";

import { createClient } from "@connectrpc/connect";
import { createConnectTransport } from "@connectrpc/connect-web";

const transport = createConnectTransport({
  baseUrl: "http://localhost:3000",
});

export const greetClient = createClient(GreetService, transport);
