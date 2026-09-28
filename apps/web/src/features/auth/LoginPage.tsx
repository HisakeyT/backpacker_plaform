import { Box, Button, TextField, Typography } from "@mui/material";

import { PageContainer } from "../../components/layout/PageContainer";

export function LoginPage() {
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
          fullWidth
        />

        <TextField
          label="Password"
          type="password"
          fullWidth
        />

        <Button
          variant="contained"
          type="button"
          fullWidth
        >
          Login
        </Button>
      </Box>
    </PageContainer>
  );
}
