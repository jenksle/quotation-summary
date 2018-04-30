﻿<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
  <meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
  <meta name="viewport" content="width=device-width, initial-scale=1" />
  <title>RJW Engineering - Quotation Report</title>

  <style type="text/css">

  @import url(https://fonts.googleapis.com/css?family=Open+Sans);

  img {
    max-width: 600px;
    outline: none;
    text-decoration: none;
    -ms-interpolation-mode: bicubic;
  }

  a {
    text-decoration: none;
    border: 0;
    outline: none;
    color: #303030;
  }

  a img {
    border: none;
  }

  td, h1, h2, h3  {
    font-family: 'Open Sans', Helvetica, Arial, sans-serif;
    font-weight: 400;
  }

  body {
    -webkit-font-smoothing:antialiased;
    -webkit-text-size-adjust:none;
    width: 100%;
    height: 100%;
    color: #f3f3f3;
    background: #f3f3f3;
    font-size: 16px;
  }

   table {
    border-collapse: collapse !important;
  }

  .headline {
    color: #ffffff;
    font-size: 36px;
  }

 .force-full-width {
	width: 100% !important;
 }

  .col-1 {
  	font-family: 'Open Sans', Helvetica, Arial, sans-serif;
  	font-size: 16px;
  	font-weight: 400;
  	line-height: 24px;
   	padding: 2px 0;
  	color: gray;
    width: 30%;
  }

  .col-2 {
  	font-family: 'Open Sans', Helvetica, Arial, sans-serif;
    font-size: 16px;
    font-weight: 400;
    line-height: 24px;
    padding: 2px 25px;
    color: #000;
    width: 70%
  }

  </style>

  <style type="text/css" media="screen">
      @media screen {
        td, h1, h2, h3 {
          font-family: 'Open Sans', 'Helvetica Neue', 'Arial', 'sans-serif' !important;
        }
      }
  </style>

  <style type="text/css" media="only screen and (max-width: 480px)">
    @media only screen and (max-width: 480px) {
      table[class="w320"] {
        width: 320px !important;
      }
    }
  </style>
</head>
<body class="body" style="padding:0; margin:0; display:block; background:#f3f3f3; -webkit-text-size-adjust:none" bgcolor="#f3f3f3">
  <table align="center" cellpadding="0" cellspacing="0" width="100%" height="100%" >
    <tr>
      <td align="center" valign="top" bgcolor="#f3f3f3"  width="100%">
        <center>
          <table style="margin: 0 auto;" cellpadding="0" cellspacing="0" width="600" class="w320">
            <tr>
              <td align="center" valign="top">

                  <!-- Pre-Header -->
                  <table style="margin: 0 auto;" cellpadding="0" cellspacing="0" width="100%">
                    <tr>
                      <td style="font-size: 13px; text-align:center; color:gray">
                        <br>
                          RJW Engineering - Quotation Summary for {{.Date}}
                        <br>
                        <br>
                      </td>
                    </tr>
                  </table>

                  <!-- Header -->
                  <table style="margin: 0 auto;" cellpadding="0" cellspacing="0" width="0%" bgcolor="#fff">
                    <tr>
                      <!-- Logo + title -->
                      <td style="padding-left:5%;padding-top:5%;padding-bottom:5%">
                        <img src="http://assetcare.rjweng.com/Content/Images/rjw-quotation-summary.jpeg" width="250" />
                      </td>
                    </tr>
                    <tr>
                      <td>
                        <table style="margin: 0 auto;" cellpadding="0" cellspacing="0" width="90%">
                          <tr>
                            <td style="font-size:16px;font-family:'Open Sans', 'Helvetica Neue', 'Arial', sans-serif !important;color:#1251A2;padding-bottom:5%">
                             Dear {{.Name}},<br><br>The table below shows a list of quotations produced on <b>{{.Date}}</b> for sales {{.AreaName}}. Quotations produced for areas not in this list will not be included in the report.
                            </td>
                          </tr>
                        </table>
                      </td>
                    </tr>
                  </table>

                  <table style="margin: 0 auto;" cellpadding="0" cellspacing="0" width="100%" bgcolor="#fff">
                    <tr>
                      <td>
                        {{if eq (len .Quotes) 0}}
                        <div style="height:10px;width:90%;margin:0 auto;border-top:1px solid gray"></div>
                        <table valign="top" border="0" cellpadding="0" cellspacing="0" width="90%" style="margin: 0 auto;">
                        <tr><td class="col-1"> No quotes produced for your area(s) today</td></tr>
                        </table>
                        {{else}}
                        <!-- Start of Block -->
                        {{range .Quotes}}

                        <div style="height:10px;width:90%;margin:0 auto;border-top:1px solid gray"></div>
                        <table valign="top" border="0" cellpadding="0" cellspacing="0" width="90%" style="margin: 0 auto;">
                          <tr>
                            <td class="col-1">
                              Quote No.
                            </td>
                            <td class="col-2">
                              <b>{{.QuoteID}}</b>
                            </td>
                          </tr>
                          <tr>
                            <td class="col-1">
                              Department
                            </td>
                            <td class="col-2">
                              {{.DepartmentName}}
                            </td>

                          </tr>
                          <tr>
                            <td class="col-1">
                              Customer
                            </td>
                            <td class="col-2">
                              {{.CustomerName}}
                            </td>
                          </tr>
                          <tr>
                            <td class="col-1">
                              Value
                            </td>
                            <td class="col-2">
                              £{{.QuoteValue}}
                            </td>
                          </tr>
                          <tr>
                            <td class="col-1">
                              Job No.
                            </td>
                            <td class="col-2">
                              {{.JobNo}}
                            </td>
                          </tr>
                          <tr>
                            <td class="col-1">
                              Description
                            </td>
                            <td class="col-2">
                              {{.Description}}
                            </td>
                          </tr>
                          <tr>
                            <td class="col-1">
                              Despatch Date
                            </td>
                            <td class="col-2">
                              {{.DateDespatched}}
                            </td>
                          </tr>
                        </table>
                        {{end}}
                        <!-- End of Block -->
                        {{end}}

            

  

                        <!-- Body Footer -->
                        <center>
                          <table style="margin: 10px auto 0 auto;" cellpadding="0" cellspacing="0" width="90%">
                            <tr>
                              <td style="font-size:16px;font-family:'Open Sans', 'Helvetica Neue', 'Arial', sans-serif !important;color:#1251A2;padding-top:20px;padding-bottom:30px;border-top:1px solid gray">
                              If you believe the quote listing shown to be incorrect, or incomplete, please email <b><a style="color:#1251A2" href="mailto:fraser@rjweng.com">fraser@rjweng.com</a></b> or call <b>07973&nbsp;692383</b>.<br><br><i>Please do not reply to this automatically generated email, as responses are not monitored.</i><br><br>
                              Best Regards,<br>
                              <b>Rewinds and J Windsor &amp; Sons (Engineers) Ltd</b>
                              </td>
                            </tr>
                          </table>
                        </center>

                      </td>
                    </tr>
                    <tr>
                      <td>
                        <div><!--[if mso]>
                          <v:roundrect xmlns:v="urn:schemas-microsoft-com:vml" xmlns:w="urn:schemas-microsoft-com:office:word" href="http://" style="height:50px;v-text-anchor:middle;width:200px;" arcsize="8%" stroke="f" fillcolor="#ac4d2f">
                            <w:anchorlock/>
                            <center>
                          <![endif]-->
                          <!--[if mso]>
                            </center>
                          </v:roundrect>
                        <![endif]--></div>
                        <br>
                        <br>
                      </td>
                    </tr>
                  </table>

                  <!-- Footer -->
                  <table style="margin: 0 auto;" cellpadding="20px" cellspacing="0" class="force-full-width" bgcolor="#f3f3f3" style="margin: 0 auto;">
                    <tr>
                      <td style="color:gray; font-size:12px; text-align: center">
                        81 Regent Road, Liverpool, Merseyside, L5 9SY<br>
                        Phone: 0151 207 2074 | Email: <a href="mailto:assetcare@rjeng.com">assetcare@rjweng.com</a><br>
                        Web: <a href="http://www.rjweng.com">www.rjweng.com</a> | Twitter: <a href="https://twitter.com/rjwengineering">@RJWEngineering</a><br>
                        Copyright © 2017 Rewinds and J. Windsor &amp; Sons (Engineers) Limited
                      </td>
                    </tr>
                  </table>

              </td>
            </tr>
          </table>
      </center>
      </td>
    </tr>
  </table>
</body>
</html>