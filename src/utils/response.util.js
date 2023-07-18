const response = (
  message,
  status,
  res,
  payload,
  count,
  prev,
  next,
  current
) => {
  res.status(status).json([
    {
      status,
      message,
      payload,
      metadata: { count, prev, next, current },
    },
  ]);
};

const serverErrorResponse = (error, res) => {
  res.status(500).json([
    {
      status: 500,
      message: "Internal server error",
      payload: error,
      metadata: { count: null, prev: null, next: null, current: null },
    },
  ]);
};

module.exports = {
  response,
  serverErrorResponse,
};
