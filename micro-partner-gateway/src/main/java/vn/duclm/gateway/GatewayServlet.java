package vn.duclm.gateway;

import com.fasterxml.jackson.databind.ObjectMapper;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import vn.duclm.gateway.entity.PaymentGatewayRes;

import javax.servlet.ServletException;
import javax.servlet.annotation.WebServlet;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import java.io.IOException;
import java.io.PrintWriter;

@WebServlet(name = "gateway", urlPatterns = "/partner", loadOnStartup = 1)
public class GatewayServlet extends HttpServlet {

    private static final Logger logger
            = LoggerFactory.getLogger(GatewayServlet.class);

    private ObjectMapper mapper = new ObjectMapper();

    @Override
    protected void doPost(HttpServletRequest req, HttpServletResponse resp) throws ServletException, IOException {
        String fnc = req.getParameter("fnc");

        PaymentGatewayRes res = new PaymentGatewayRes();
        res.setFnc(fnc);
        res.setPgUserCode("EVNNPC_VIPOST");
        res.setData("Sample data");

        String resStr = mapper.writeValueAsString(res);

        PrintWriter out = resp.getWriter();
        resp.setContentType("application/json");
        resp.setCharacterEncoding("UTF-8");
        out.print(resStr);
        out.flush();

        logger.info("Receive function: fnc={}", fnc);
    }
}
