import { Box, Typography } from "@mui/material";

export function Footer() {
  return (
    <Box
      className="Footer"
      component="footer"
      sx={{
        py: 3,
        textAlign: "center",
      }}
    >
      <Typography variant="body2" color="text.secondary">
        © Backpacker
      </Typography>
    </Box>
  );
}
