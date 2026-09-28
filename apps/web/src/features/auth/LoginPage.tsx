import { useState } from "react";
import { Box, Button, TextField, Typography } from "@mui/material";
import { PageContainer } from "../../components/layout/PageContainer";
import { login } from "./api";

export function LoginPage() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  const handleSubmit = async () => {
    try {
      const result = await login({
        email,
        password,
      });

      console.log(result);
      localStorage.setItem("token", result.token);
    } catch (error) {
      console.error(error);
    }
  };

  return (
    <PageContainer>
      <Box
        sx={{
          maxWidth: 480,
          mx: "auto",
          display: "flex",
          flexDirection: "column",
          gap: 2,
        }}
      >
        <Typography variant="h4" component="h1">
          Login
        </Typography>

        <TextField
          label="Email"
          type="email"
          value={email}
          onChange={(event) => setEmail(event.target.value)}
          fullWidth
        />

        <TextField
          label="Password"
          type="password"
          value={password}
          onChange={(event) => setPassword(event.target.value)}
          fullWidth
        />

        <Button
          variant="contained"
          type="button"
          fullWidth
          onClick={handleSubmit}
        >
          Login
        </Button>
      </Box>
    </PageContainer>
  );
}
