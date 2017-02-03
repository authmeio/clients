/**
 * Created by parth on 2/2/17.
 */

    // @include jquery.js
    // @include sha265.js


var Authme = function (apiKey, apiSecret) {
        return {
            "getLoginToken": function (email, referenceId, callback) {
                var hashString = [apiKey, email, apiSecret].join("|");
                var hash = Sha265.hash(hashString);
                var l = function (email, referenceId, callback, tries) {
                    if (tries < 0) {
                        callback({
                            Status: "failed",
                            ReferenceId: ""
                        });
                        return
                    }
                    $.ajax({
                        url: 'https://authme.io/v1/trylogin',
                        data: JSON.stringify({
                            Email: email,
                            ReferenceId: referenceId,
                            Hash: hash
                        }),
                        contentType: "application/json",
                        headers: {
                            "X-Api-Key": apiKey
                        },
                        method: "POST",
                        success: function (response) {
                            if (response.Status == "initiated" || response.Status == "pending" || response.Status == "auth_initiated") {
                                setTimeout(function () {
                                    l(email, response.ReferenceId, callback, tries - 1)
                                }, 1500)
                            }
                            if (response.Status == "authorized") {
                                callback({
                                    Status: "success",
                                    ReferenceId: response.ReferenceId
                                })
                            } else if (response.Status == "rejected") {
                                callback({
                                    Status: "failed",
                                    ReferenceId: response.ReferenceId
                                })

                            }
                            //console.log(response);
                        },
                        error: function () {
                            alert("Login failed")
                        }
                    })
                };
                l(email, referenceId, callback, 15)
            }
        }
    };
