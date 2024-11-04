FROM alpine:latest
LABEL maintainer="Maher Sidawy <msidawy@africell.com>"
LABEL decsription="IN live feed consumer Activation Promo"
RUN mkdir /kafka
COPY /afr_DRC_INLiveFeed_consumer_SSR_ActivationPromoEast_d /afr_DRC_INLiveFeed_consumer_SSR_ActivationPromoEast_d
#WORKDIR /brmgw
EXPOSE 9662/tcp
ENTRYPOINT ["/afr_DRC_INLiveFeed_consumer_SSR_ActivationPromoEast_d"]
CMD [ "/afr_DRC_INLiveFeed_consumer_SSR_ActivationPromoEast_d" ]
