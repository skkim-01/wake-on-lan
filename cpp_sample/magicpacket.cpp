#include "StdAfx.h"
#include "WOL/WakeOnLan.h"
#include "BroadCast/BroadcastSender.h"

#define MAGIC_PACKET_SIZE	102

CWakeOnLan::CWakeOnLan(void)
{
}


CWakeOnLan::~CWakeOnLan(void)
{
}


BOOL CWakeOnLan::GetMagicPacket(char* szMacAddress)
{
	char magicp[MAGIC_PACKET_SIZE] = {0, };
	char buff[3] = {0, };
	
	int nMacLen = strlen(szMacAddress);
	BOOL bToken = TRUE;
	if( nMacLen == 17 )
	{
		bToken = TRUE;
	}
	else if( nMacLen == 12 )
	{
		bToken = FALSE;
	}
	else
	{
		return FALSE;
	}


    for( int i = 0; i < 6; i++ )
	{
        magicp[i] = 0xff;
	}

	if( bToken )
	{
		for( int i = 0; i < 6; i ++ )
		{
			memcpy( buff, szMacAddress + ( i * 3 ), 2 );
			magicp[ i + 6 ] = HexStr2Int( buff );
		}
	}
	else
	{
		for( int i = 0; i < 6; i ++ )
		{
			memcpy( buff, szMacAddress + ( i * 2 ), 2 );
			magicp[ i + 6 ] = HexStr2Int( buff );
		}
	}

    for( int i = 0; i < 15; i++ )
    {
        memcpy( &magicp[ ( i + 2 ) * 6 ], &magicp[6], 6 );
    }
}


int CWakeOnLan::HexStr2Int( char* szHex )
{
    if( 0 == szHex || 2 < strlen( szHex ) )
	{
        return 0;
	}

    int res = 0;
    char* stop = 0;

    res = strtol( szHex, &stop, 16 );

    if( LONG_MAX == res || LONG_MIN == res || 0 == res )
    {
        return 0;
    }

    return res;
}